package http

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/model"
	"github.com/lfxnxf/school/api-gateway/utils"
	"go.uber.org/zap"
	"unicode/utf8"
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

func editUserInfo(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "editUserInfo")

	var (
		req model.EditUserInfoReq
	)

	if err := school_http.Requests.Body(c.Ctx, c.Request).ParseJson(&req).Error(); err != nil {
		log.Errorw("Parse Param",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	atom, err := school_http.Requests.Query(c.Ctx, c.Request).Parse(&req).Atom()
	if err != nil {
		log.Errorw("parse body error", zap.String("err", err.Error()))
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	if req.Name == "" {
		log.Errorw("name is empty error", zap.String("err", err.Error()))
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	if utf8.RuneCountInString(req.Name) > 10 || !utils.IsAllChinese(req.Name) {
		log.Errorw("name length > 10 or name is not chinese", zap.String("err", err.Error()))
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	resp, err := svc.EditUserInfo(c.Ctx, atom, req)
	if err != nil {
		log.Errorw("EditUserInfo",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, err)
		return
	}

	log.Infow("success")

	c.JSON(resp, nil)
}