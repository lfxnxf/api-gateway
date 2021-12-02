package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/conf"
	"github.com/lfxnxf/school/api-gateway/dao"
	"github.com/lfxnxf/school/api-gateway/manager"
	"github.com/lfxnxf/school/api-gateway/utils"
	"go.uber.org/zap"
	"sync"
)

const (
	allUidTopic    = ""
	eventSubscribe = "subscribe" // 订阅类型
	eventMessage   = "message"
)

type Ws struct {
	c *conf.Config

	// dao: db handler
	dao *dao.Dao

	// manager: other client(s), other middleware(s)
	mgr *manager.Manager
}

func New(c *conf.Config) *Ws {
	return &Ws{
		c:   c,
		dao: dao.New(c),
		mgr: manager.New(c),
	}
}

// ClientManager is a websocket manager
type ClientManager struct {
	Clients        map[string][]*Client
	ClientTopicMap *sync.Map
	broadcast      chan DownsideMessage
	register       chan *Client
	unregister     chan *Client
}

// 下行消息
type DownsideMessage struct {
	Sender    int64       `json:"sender"`
	Topic     string      `json:"topic"`     // 主题
	Recipient int64       `json:"recipient"` // 指定接收人
	Content   interface{} `json:"content"`
	Event     string      `json:"event"`
	SeqId     int64       `json:"seq_id"`
	ErrorCode int64       `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
}

// 上行消息
type UpsideMessage struct {
	Msg       []byte `json:"msg"`
	Event     string `json:"event"`
	SeqId     int64  `json:"seq_id"`
	Namespace string `json:"namespace"`
	Topic     string `json:"topic"`
}

// Client is a websocket client
type Client struct {
	UID       int64
	Socket    *websocket.Conn
	Send      chan []byte
	TopicList []string
}

// Manager define a ws server manager
var Manager = ClientManager{
	broadcast:      make(chan DownsideMessage),
	register:       make(chan *Client),
	unregister:     make(chan *Client),
	Clients:        make(map[string][]*Client),
	ClientTopicMap: new(sync.Map),
}

func (m *ClientManager) Register(client *Client) {
	m.register <- client
}

// Start is to start a ws server
func (m *ClientManager) Start() {
	for {
		select {
		// 注册用户
		case conn := <-m.register:
			// 默认订阅单个频道
			topic := topicKey("uid:", conn.UID)
			m.Clients[topic][0] = conn
			m.ClientTopicMap.Store(topic, conn)
			conn.TopicList = append(conn.TopicList, topic)

		// 用户退出长连接
		case conn := <-m.unregister:
			for _, topic := range conn.TopicList {
				for i := 0; i < len(m.Clients[topic]); i++ {
					if conn != m.Clients[topic][i] {
						continue
					}
					m.Clients[topic] = append(m.Clients[topic][:i], m.Clients[topic][i:]...)
				}
				m.ClientTopicMap.Delete(topic)
			}
		case sendMs := <-m.broadcast:
			var sendConnList []*Client
			if sendMs.Recipient != 0 {
				sendConnLoad, _ := m.ClientTopicMap.Load(topicKey(sendMs.Topic, sendMs.Recipient))
				sendConn, ok := sendConnLoad.(*Client)
				if !ok {
					break
				}
				sendConnList = append(sendConnList, sendConn)
			} else {
				for _, v := range m.Clients[sendMs.Topic] {
					sendConnList = append(sendConnList, v)
				}
			}
			if len(sendConnList) <= 0 {
				break
			}
			bytes, _ := json.Marshal(sendMs)

			for _, conn := range sendConnList {
				select {
				case conn.Send <- bytes:
				default:
					close(conn.Send)
					m.unregister <- conn
				}
			}
		}
	}
}

func topicKey(topic string, uid int64) string {
	return fmt.Sprintf("%s_%d", topic, uid)
}

func (m *ClientManager) SendOne(data DownsideMessage) {
	select {
	case m.broadcast <- data:
		break
	default:
	}
}

func (m *ClientManager) MultiSend(recipientList []int64, msg interface{}, ev, topic string) {
	ms, _ := json.Marshal(msg)
	for _, v := range recipientList {
		select {
		case m.broadcast <- DownsideMessage{
			Sender:    0,
			Topic:     topic,
			Recipient: v,
			Content:   ms,
			Event:     ev,
		}:
			break
		default:

		}

	}
}

func (m *ClientManager) Broadcast(topic, ev string, msg interface{}) {
	ms, _ := json.Marshal(msg)
	select {
	case m.broadcast <- DownsideMessage{
		Sender:    0,
		Topic:     topic,
		Recipient: 0,
		Content:   ms,
		Event:     ev,
	}:
		break
	default:
	}
}

func (c *Client) Read(ctx context.Context) {
	log := logging.For(ctx, "func", "ws.Read")
	defer func() {
		Manager.unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.unregister <- c
			_ = c.Socket.Close()
			break
		}
		if string(message) == "ping" {
			Manager.SendOne(DownsideMessage{
				Sender:    0,
				Topic:     "",
				Recipient: c.UID,
				Content:   "pong",
				Event:     "",
			})
			continue
		}
		log.Infow("ws read message:", zap.String("data", fmt.Sprintf("uid:%d, data:%s", c.UID, string(message))))
		var upsideMessage UpsideMessage
		err = json.Unmarshal(message, &upsideMessage)
		if err != nil {
			Manager.SendOne(DownsideMessage{
				Sender:    0,
				Topic:     "",
				Recipient: c.UID,
				Content:   nil,
				Event:     "",
				ErrorCode: 499,
				ErrorMsg:  "参数错误",
			})
			continue
		}
		// 订阅
		if upsideMessage.Event == eventSubscribe {
			topic := upsideMessage.Topic
			for i := 0; i < len(Manager.Clients[topic]); i++ {
				if Manager.Clients[topic][i].UID != c.UID {
					continue
				}
				Manager.Clients[topic] = append(Manager.Clients[topic][:i], Manager.Clients[topic][i:]...)
			}
			Manager.Clients[topic] = append(Manager.Clients[topic], c)
			if !utils.InStringArray(topic, c.TopicList) {
				c.TopicList = append(c.TopicList, topic)
			}
			continue
		}
		//TODO 调用业务数据
		jsonMessage := DownsideMessage{}
		Manager.broadcast <- jsonMessage
	}
}

func (c *Client) Write(ctx context.Context) {
	log := logging.For(ctx, "func", "ws.Write")
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			log.Infow("ws send message:", zap.String("data", fmt.Sprintf("uid:%d, data:%s", c.UID, string(message))))
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
