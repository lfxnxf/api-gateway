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
	AddDriverLockKey         = "school:add:driver:%d"
	AddDriverLockKeyExpire   = 1
	AddDriverLockKeyLoopTime = 1
)

// 获取全部司机信息
func (s *Service) GetDivers(ctx context.Context, atom *school_http.Atom) (interface{}, error) {
	log := logging.For(ctx, "func", "GetDivers",
		zap.Int64("uid", atom.Uid),
	)

	drivers, err := s.dao.GetDrivers(ctx, atom.Uid)
	if err != nil {
		log.Errorw("s.dao.GetAllVehicleByBoss error", zap.Error(err))
		return nil, err
	}
	resp := model.GetDriversResp{
		List: make([]model.DriverInfo, 0),
	}
	for _, v := range drivers {
		info := model.DriverInfo{
			Id:       v.Id,
			Name:     v.Name,
			Phone:    v.Phone,
			Identity: v.Identity,
		}
		resp.List = append(resp.List, info)
	}

	log.Infow("success!")
	return resp, nil
}

// 新增司机
func (s *Service) AddDiver(ctx context.Context, atom *school_http.Atom, req model.AddDriverReq) (interface{}, error) {
	log := logging.For(ctx, "func", "AddDiver",
		zap.Int64("uid", atom.Uid),
		zap.Any("req", req),
	)

	// 分布式锁
	dispersedLock := dispersed_lock.New(ctx, dao.RedisClient, fmt.Sprintf(AddDriverLockKey, req.Phone), AddDriverLockKeyExpire)
	dispersedLock.LoopLock(ctx, AddDriverLockKeyLoopTime)
	defer dispersedLock.Unlock(ctx)

	// 判断用户是否有添加用户权限
	bossUser, err := s.dao.GetUserInfo(ctx, atom.Uid)
	if err != nil {
		log.Errorw("s.dao.GetUserInfo error", zap.Error(err))
		return nil, err
	}

	// 添加管理员
	if bossUser.Identity != model.IdentityBoss && req.Identity == model.IdentityAdmin {
		return nil, error_code.HasNotAuthAddAdmin
	}

	// 添加司机
	if req.Identity == model.IdentityDriver && bossUser.Identity != model.IdentityBoss && bossUser.Identity != model.IdentityAdmin {
		return nil, error_code.HasNotAuthAddDriver
	}

	// 判断手机号是否已存在
	user, err := s.dao.GetUserInfoByPhone(ctx, req.Phone)
	if err != nil {
		log.Errorw("s.dao.GetAllVehicleByBoss error", zap.Error(err))
		return nil, err
	}

	if user.Id > 0 {
		return nil, error_code.HasEqPhone
	}

	err = s.dao.InsertUsers(ctx, model.UsersModel{
		Name:     req.Name,
		BossId:   atom.Uid,
		Identity: req.Identity,
		Phone:    req.Phone,
	})
	if err != nil {
		log.Errorw("s.dao.InsertUsers error", zap.Error(err))
		return nil, err
	}

	log.Infow("success!")
	return nil, nil
}
