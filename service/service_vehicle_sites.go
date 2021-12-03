package service

import (
	"context"
	"fmt"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/dao"
	"github.com/lfxnxf/school/api-gateway/dispersed_lock"
	"github.com/lfxnxf/school/api-gateway/error_code"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

const (
	AddVehicleSitesLockKey         = "school:add:vehicle:sites:%d"
	AddVehicleSitesLockKeyExpire   = 1
	AddVehicleSitesLockKeyLoopTime = 1
)

// 保存车辆站点
func (s *Service) SaveVehicleSites(ctx context.Context, atom *school_http.Atom, req model.SaveVehicleSitesReq) (interface{}, error) {
	log := logging.For(ctx, "func", "AddVehicleSites",
		zap.Int64("uid", atom.Uid),
		zap.Any("req", req),
	)

	// 分布式锁
	dispersedLock := dispersed_lock.New(ctx, dao.RedisClient, fmt.Sprintf(AddVehicleSitesLockKey, req.VehicleId), AddVehicleSitesLockKeyExpire)
	dispersedLock.LoopLock(ctx, AddVehicleSitesLockKeyLoopTime)
	defer dispersedLock.Unlock(ctx)

	var err error

	// 获取车辆信息
	vehicleInfo, err := s.dao.GetVehicleInfoById(ctx, req.VehicleId)
	if err != nil {
		log.Errorw("s.dao.GetVehicleInfoById, err:", zap.String("error", err.Error()))
		return nil, err
	}

	// 不是老板，没权限
	if vehicleInfo.BossId != atom.Uid {
		return nil, error_code.NotHasAuth
	}

	// 先删除车辆原本站点
	tx := s.dao.StartTransaction(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	err = s.dao.DeleteVehicleSites(ctx, tx, req.VehicleId)
	if err != nil {
		log.Errorw("s.dao.DeleteVehicleSites, err:", zap.String("error", err.Error()))
		return nil, err
	}

	// 新增车辆站点
	var modelList []model.VehicleSitesModel
	for _, v := range req.List {
		m := model.VehicleSitesModel{
			VehicleId: req.VehicleId,
			SiteId:    v.SiteId,
			Sort:      v.Sort,
		}
		modelList = append(modelList, m)
	}
	err = s.dao.InsertVehicleSites(ctx, tx, modelList)
	if err != nil {
		log.Errorw("s.dao.InsertVehicleSites, err:", zap.String("error", err.Error()))
		return nil, err
	}

	log.Infow("success!")
	return nil, nil
}
