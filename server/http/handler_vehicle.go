package http

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

type DefaultShifts struct {
	Name string `json:"name"`
}
var defaultShifts = []DefaultShifts{
	DefaultShifts{Name: "早班（上学）"},
	DefaultShifts{Name: "中班（放学）"},
	DefaultShifts{Name: "中班（上学）"},
	DefaultShifts{Name: "晚班（放学）"},
}

func saveVehicle(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "saveVehicle")

	var (
		req model.SaveVehicleReq
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

	if req.LicensePlate == "" {
		log.Errorw("req.LicensePlate empty")
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	if req.DriverId == 0 {
		log.Errorw("req.DriverId empty")
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	if len(req.ShiftsList) <= 0 {
		log.Errorw("req.ShiftsList empty")
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	if len(req.VehicleSiteList) <= 0 {
		log.Errorw("req.VehicleSiteList empty")
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	resp, err := svc.SaveVehicle(c.Ctx, atom, req)
	if err != nil {
		log.Errorw("SaveVehicle",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, err)
		return
	}

	log.Infow("success")

	c.JSON(resp, nil)
}

func getAllVehicleInfo(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "getAllVehicleInfo")

	atom, err := school_http.Requests.Query(c.Ctx, c.Request).Atom()
	if err != nil {
		log.Errorw("parse body error", zap.String("err", err.Error()))
		c.JSONAbort(nil, school_errors.Codes.ClientError)
		return
	}

	resp, err := svc.GetAllVehicleInfo(c.Ctx, atom)
	if err != nil {
		log.Errorw("GetAllVehicleInfo",
			zap.String("err", err.Error()),
		)
		c.JSONAbort(nil, err)
		return
	}

	log.Infow("success")

	c.JSON(resp, nil)
}

func getDefaultShifts(c *httpserver.Context) {
	c.JSON(map[string][]DefaultShifts{
		"list": defaultShifts,
	}, nil)
}
