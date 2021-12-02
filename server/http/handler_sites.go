package http

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

// 获取站点列表
func getSitesList(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "getSitesList")

	var req model.GetSitesListReq

	atom, err := school_http.Requests.Query(c.Ctx, c.Request).Parse(&req).Atom()
	if err != nil {
		log.Errorw("parse body error", zap.String("err", err.Error()))
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	resp, err := svc.GetSitesList(c.Ctx, atom, req)
	if err != nil {
		log.Errorw("GetSitesList",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, err)
		return
	}

	log.Infow("success")

	c.JSON(resp, nil)
}

