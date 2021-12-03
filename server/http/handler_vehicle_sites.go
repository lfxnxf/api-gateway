package http

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

func saveVehicleSites(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "saveVehicleSites")

	var (
		req model.SaveVehicleSitesReq
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

	if req.VehicleId <= 0 {
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	if len(req.List) <= 0 {
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	resp, err := svc.SaveVehicleSites(c.Ctx, atom, req)
	if err != nil {
		log.Errorw("AddVehicle",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, err)
		return
	}

	log.Infow("success")

	c.JSON(resp, nil)
}
