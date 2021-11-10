package http

import (
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/school/api-gateway/ws"
)

type Request struct {
	AppName     string `json:"app_name"`
	ServiceName string `json:"service_name"`
}

var pathMap = map[string]Request{
	"/api/123/456/5778": {
		AppName:     "socialgame",
		ServiceName: "game.center.logic",
	},
}

func ping(c *httpserver.Context) {
	if err := svc.Ping(c.Ctx); err != nil {
		c.JSONAbort(nil, err)
		return
	}
	okMsg := map[string]string{"result": "ok"}
	c.JSON(okMsg, nil)
}

func actionHandler(c *httpserver.Context) {
	//action := c.Params.ByName("action")
	ws.Manager.SendOne(ws.DownsideMessage{
		Sender:    0,
		Topic:     "",
		Recipient: 1310265,
		Content:   c.Request.URL.RawQuery,
		Event:     "",
		SeqId:     1,
		ErrorCode: 0,
		ErrorMsg:  "",
	})
	c.JSON(struct {
		Message string `json:"message"`
	}{
		Message: c.Request.URL.RawQuery,
	}, nil)
}
