package service

import (
	"context"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
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

