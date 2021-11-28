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
	"go.uber.org/zap"
	"sync"
)

const (
	allUidTopic = ""
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
}

// Client is a websocket client
type Client struct {
	UID    int64
	Socket *websocket.Conn
	Send   chan []byte
	Topic  string `json:"topic"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	broadcast:      make(chan DownsideMessage),
	register:       make(chan *Client),
	unregister:     make(chan *Client),
	Clients:        make(map[string][]*Client),
	ClientTopicMap: new(sync.Map),
}

func (c *ClientManager) Register(client *Client) {
	c.register <- client
}

// Start is to start a ws server
func (c *ClientManager) Start() {
	for {
		select {
		// 注册用户，存入对应topic数组中
		case conn := <-c.register:
			key := topicKey(conn.Topic, conn.UID)
			_, ok := c.ClientTopicMap.Load(key)
			if ok {
				c.ClientTopicMap.Delete(key)
				for i := 0; i < len(c.Clients[conn.Topic]); i++ {
					if c.Clients[conn.Topic][i].UID != conn.UID {
						continue
					}
					c.Clients[conn.Topic] = append(c.Clients[conn.Topic][:i], c.Clients[conn.Topic][i:]...)
				}
			}
			c.Clients[conn.Topic] = append(c.Clients[conn.Topic], conn)
			c.ClientTopicMap.Store(key, conn)

		// 用户退出长连接
		case conn := <-c.unregister:
			for i := 0; i < len(c.Clients[conn.Topic]); i++ {
				if conn != c.Clients[conn.Topic][i] {
					continue
				}
				c.Clients[conn.Topic] = append(c.Clients[conn.Topic][:i], c.Clients[conn.Topic][i:]...)
			}
			c.ClientTopicMap.Delete(topicKey(conn.Topic, conn.UID))
		case sendMs := <-c.broadcast:
			var sendConnList []*Client
			if sendMs.Recipient != 0 {
				sendConnLoad, _ := c.ClientTopicMap.Load(topicKey(sendMs.Topic, sendMs.Recipient))
				sendConn, ok := sendConnLoad.(*Client)
				if !ok {
					break
				}
				sendConnList = append(sendConnList, sendConn)
			} else {
				for _, v := range c.Clients[sendMs.Topic] {
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
					c.unregister <- conn
				}
			}
		}
	}
}

func topicKey(topic string, uid int64) string {
	return fmt.Sprintf("%s_%d", topic, uid)
}

func (c *ClientManager) SendOne(data DownsideMessage) {
	select {
	case c.broadcast <- data:
		break
	default:
	}
}

func (c *ClientManager) MultiSend(recipientList []int64, msg interface{}, ev, topic string) {
	ms, _ := json.Marshal(msg)
	for _, v := range recipientList {
		select {
		case c.broadcast <- DownsideMessage{
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

func (c *ClientManager) Broadcast(topic, ev string, msg interface{}) {
	ms, _ := json.Marshal(msg)
	select {
	case c.broadcast <- DownsideMessage{
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
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.unregister <- c
			c.Socket.Close()
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
