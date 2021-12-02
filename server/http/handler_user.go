package http

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"go.uber.org/zap"
)

func getUserInfo(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "getUserInfo")

	atom, err := school_http.Requests.Query(c.Ctx, c.Request).Atom()
	if err != nil {
		log.Errorw("parse body error", zap.String("err", err.Error()))
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	resp, err := svc.GetUserInfo(c.Ctx, atom)
	if err != nil {
		log.Errorw("GetUserInfo",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, err)
		return
	}

	log.Infow("success")

	c.JSON(resp, nil)
}

