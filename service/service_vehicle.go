package service

import (
	"fmt"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/dao"
	"github.com/lfxnxf/school/api-gateway/dispersed_lock"
	"github.com/lfxnxf/school/api-gateway/error_code"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

const (
	AddVehicleLockKey         = "school:add:vehicle:%d"
	AddVehicleLockKeyExpire   = 1
	AddVehicleLockKeyLoopTime = 1
)

// 新增车辆
func (s *Service) AddVehicle(ctx context.Context, atom *school_http.Atom, req model.AddVehicleReq) (interface{}, error) {
	log := logging.For(ctx, "func", "AddVehicle",
		zap.Any("req", req),
	)

	// 分布式锁
	dispersedLock := dispersed_lock.New(ctx, dao.RedisClient, fmt.Sprintf(AddVehicleLockKey, atom.Uid), AddVehicleLockKeyExpire)
	dispersedLock.LoopLock(ctx, AddVehicleLockKeyLoopTime)
	defer dispersedLock.Unlock(ctx)

	// 判断车牌是否重复
	vehicleInfo, err := s.dao.GetVehicleInfoByLicensePlate(ctx, req.LicensePlate)
	if err != nil {
		log.Errorw("s.dao.GetVehicleInfoByLicensePlate, err:", zap.String("error", err.Error()))
		return nil, err
	}

	// 已有相同车牌号
	if vehicleInfo.Id > 0 {
		return nil, error_code.HasEqLicensePlate
	}

	// 新增车辆信息
	err = s.dao.InsertVehicle(ctx, model.VehicleModel{
		BossId:        atom.Uid,
		VehicleInfoId: req.VehicleInfoId,
		LicensePlate:  req.LicensePlate,
		DriverId:      req.DriverId,
	})
	if err != nil {
		log.Errorw("s.dao.InsertVehicle, err:", zap.String("error", err.Error()))
		return nil, err
	}

	log.Infow("success!")
	return nil, nil
}

func (s *Service) GetAllVehicleInfo(ctx context.Context, atom *school_http.Atom) (interface{}, error) {
	log := logging.For(ctx, "func", "GetAllVehicleInfo",
		zap.Int64("uid", atom.Uid),
	)

	VehicleList, err := s.dao.GetAllVehicleInfo(ctx)
	if err != nil {
		log.Errorw("s.dao.GetAllVehicleInfo, err:", zap.String("error", err.Error()))
		return nil, err
	}

	resp := model.GetVehicleInfoResp{
		List: make([]model.VehicleInfo, 0),
	}
	for _, v := range VehicleList {
		info := model.VehicleInfo{
			Id:      v.Id,
			Name:    v.Name,
			LoadNum: v.LoadNum,
		}
		resp.List = append(resp.List, info)
	}

	log.Infow("success!")
	return resp, nil
}
