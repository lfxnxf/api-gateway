package http

import (
	"fmt"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	httpplugin "github.com/lfxnxf/frame/logic/inits/plugins/http"
	"github.com/lfxnxf/school/api-gateway/conf"
	"github.com/lfxnxf/school/api-gateway/error_code"
	"github.com/lfxnxf/school/api-gateway/service"
	"github.com/lfxnxf/school/api-gateway/utils"
	"github.com/lfxnxf/school/api-gateway/ws"
	"go.uber.org/zap"
)

var (
	svc *service.Service

	wsSvc *ws.Ws

	httpServer httpserver.Server
)

// 验证token白名单
var TokenWhitePath = []string{
	"/api/login",
	"/api/login/send_verification_code",
}

// Init create a rpc server and run it
func Init(s *service.Service, w *ws.Ws, conf *conf.Config) {
	svc = s

	wsSvc = w

	// new http server
	httpServer = inits.HTTPServer()

	// add namespace plugin
	httpServer.Use(httpplugin.Namespace)

	// 验证token
	httpServer.Use(CheckToken)

	// register handler with http route
	initRoute(httpServer)

	// start a http server
	go func() {
		if err := httpServer.Run(); err != nil {
			logging.Fatalf("http server start failed, err %v", err)
		}
	}()

}

func CheckToken(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "CheckToken")
	path := c.Request.URL.Path
	if utils.InStringArray(path, TokenWhitePath) {
		return
	}
	token := c.Request.Header.Get("auth_token")
	log.Infow("get token:", zap.String("token", token))
	user, err := svc.GetUserByToken(c.Ctx, token)
	if err != nil {
		c.JSONAbort(nil, err)
		return
	}
	if user.Id <= 0 {
		c.JSONAbort(nil, error_code.UnLogin)
		return
	}
	u := "uid=%d"
	if c.Request.URL.RawQuery != "" {
		u = fmt.Sprintf("&%s", u)
	}
	c.Request.URL.RawQuery = fmt.Sprintf("%s%s", c.Request.URL.RawQuery, fmt.Sprintf(u, user.Id))

	// 刷新token
	go func() {
		err = svc.RefreshToken(c.Ctx, token)
		if err != nil {
			logging.Errorw("svc.RefreshToken error", zap.Error(err))
		}
	}()
}

func Shutdown() {
	if httpServer != nil {
		httpServer.Stop()
	}
	if svc != nil {
		svc.Close()
	}
}
