package http

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

// @Summary 新增车辆信息
// @Description 新增车辆信息
// @Accept  json
// @Produce  json
// @Param  body body model.AddVehicleReq true "新增车辆信息参数"
// @Success 200 {object} utils.WrapResp
// @Failure 500 {object} utils.WrapResp
// @Router /api/v1/vehicle/add [post]
func addVehicle(c *httpserver.Context) {
	log := logging.For(c.Ctx, "func", "addVehicle")

	var (
		req model.AddVehicleReq
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

	resp, err := svc.AddVehicle(c.Ctx, atom, req)
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
