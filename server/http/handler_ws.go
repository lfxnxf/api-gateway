package http

import (
	"github.com/gorilla/websocket"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/ws"
	"net/http"
)

func wsHandler(c *httpserver.Context) {
	// change the reqest to websocket model
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).
		Upgrade(c.Response.Writer(), c.Request, nil)

	if err != nil {
		http.NotFound(c.Response.Writer(), c.Request)
		return
	}
	atom, err := school_http.Requests.Query(c.Ctx, c.Request).Atom()
	if err != nil {
		c.JSON(nil, school_errors.Code{})
		return
	}
	// websocket connect
	client := &ws.Client{UID: atom.Uid, Socket: conn, Send: make(chan []byte)}

	go ws.Manager.Register(client)
	go client.Read(c.Ctx)
	go client.Write(c.Ctx)
}