package service

import (
	"context"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/error_code"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

// 获取用户信息
func (s *Service) GetUserInfo(ctx context.Context, atom *school_http.Atom) (interface{}, error) {
	log := logging.For(ctx, "func", "GetDivers",
		zap.Int64("uid", atom.Uid),
	)

	user, err := s.dao.GetUserInfo(ctx, atom.Uid)
	if err != nil {
		log.Errorw("s.dao.GetAllVehicleByBoss error", zap.Error(err))
		return nil, err
	}
	resp := model.GetUserInfoResp{
		Uid:      user.Id,
		Name:     user.Name,
		Address:  user.Address,
		Phone:    user.Phone,
		Identity: user.Identity,
	}

	log.Infow("success!")
	return resp, nil
}

// 修改用户信息
func (s *Service) EditUserInfo(ctx context.Context, atom *school_http.Atom, req model.EditUserInfoReq) (interface{}, error) {
	log := logging.For(ctx, "func", "GetDivers",
		zap.Int64("uid", atom.Uid),
		zap.Any("req", req),
	)

	updateMap := make(map[string]interface{})
	if req.Address != "" {
		updateMap["address"] = req.Address
	}

	if req.Name != "" {
		updateMap["name"] = req.Name
	}

	// 先获取用户身份信息
	userInfo, err := s.dao.GetUserInfo(ctx, atom.Uid)
	if err != nil {
		log.Errorw("s.dao.GetUserInfo error", zap.Error(err))
		return nil, err
	}

	if req.Phone != 0 {
		//todo 当前阶段不可以更改手机号，如果手机号不同直接报错
		if req.Phone != userInfo.Phone {
			return nil, error_code.CannotEditPhone
		}
	}

	err = s.dao.UpdateUsersByMap(ctx, atom.Uid, updateMap)
	if err != nil {
		log.Errorw("s.dao.UpdateUsersByMap error", zap.Error(err))
		return nil, err
	}

	log.Infow("success!")
	return nil, nil
}
