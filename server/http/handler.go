package http

import (
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
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
	c.JSON(struct {
		Path  string `json:"path"`
		Query string `json:"query"`
	}{
		Path:  c.Request.URL.Path,
		Query: c.Request.URL.RawQuery,
	}, nil)
}
