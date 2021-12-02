package http

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
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

func login(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "login")
	var (
		req model.LoginReq
	)

	if err := school_http.Requests.Body(c.Ctx, c.Request).ParseJson(&req).Error(); err != nil {
		log.Errorw("Parse Param",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	if req.Phone == 0 {
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	if req.VerificationCode == "" {
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	resp, err := svc.Login(c.Ctx, req)
	if err != nil {
		log.Errorw("Login",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, err)
		return
	}

	log.Infow("success")

	c.JSON(resp, nil)
}

func sendVerificationCode(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "sendVerificationCode")
	var (
		req model.SendVerificationCodeReq
	)

	if err := school_http.Requests.Body(c.Ctx, c.Request).ParseJson(&req).Error(); err != nil {
		log.Errorw("Parse Param",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	//todo 正则表达式判断手机号是否正确
	if req.Phone == 0 {
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	resp, err := svc.SendVerificationCode(c.Ctx, req)
	if err != nil {
		log.Errorw("SendVerificationCode",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, err)
		return
	}

	log.Infow("success")

	c.JSON(resp, nil)
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
