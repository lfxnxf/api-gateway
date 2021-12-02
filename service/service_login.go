package service

import (
	"fmt"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/error_code"
	"github.com/lfxnxf/school/api-gateway/model"
	"github.com/lfxnxf/school/api-gateway/utils"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// 通过token获取用户信息
func (s *Service) GetUserByToken(ctx context.Context, token string) (model.UsersModel, error) {
	log := logging.For(ctx, "func", "GetDivers",
		zap.String("token", token),
	)

	resp, err := s.dao.GetUserByToken(ctx, token)
	if err != nil {
		log.Errorw("s.dao.GetUserByToken err", zap.Error(err))
		return resp, err
	}

	log.Infow("success!")
	return resp, nil
}

// 登录
func (s *Service) Login(ctx context.Context, req model.LoginReq) (interface{}, error) {
	log := logging.For(ctx, "func", "Login",
		zap.Any("req", req),
	)

	//todo 判断验证码是否正确
	code, err := s.dao.GetVerificationCode(ctx, req.Phone)
	if code != req.VerificationCode {
		return nil, error_code.VerificationCodeWrong
	}

	// 通过手机号获取用户信息
	user, err := s.dao.GetUserInfoByPhone(ctx, req.Phone)
	if err != nil {
		log.Errorw("s.dao.GetUserInfoByPhone err", zap.Error(err))
		return nil, err
	}

	// 生成token
	randNum := utils.Random(100000, 999999)
	token := utils.Md5(fmt.Sprintf("%d%d", req.Phone, randNum))

	if user.Id == 0 {
		// 注册
		user, err = s.dao.InsertUsers(ctx, model.UsersModel{
			Identity: model.IdentityBoss,
			Phone:    req.Phone,
			Token:    token,
		})
		if err != nil {
			log.Errorw("s.dao.InsertUsers err", zap.Error(err))
			return nil, err
		}
	} else {
		// 登录
		err = s.dao.UpdateUsersByMap(ctx, user.Id, map[string]interface{}{
			"token": token,
		})
		if err != nil {
			log.Errorw("s.dao.UpdateUsersByMap err", zap.Error(err))
			return nil, err
		}
		user.Token = token
	}

	// redis中设置token
	err = s.dao.SetToken(ctx, token, user)
	if err != nil {
		log.Errorw("s.dao.SetToken err", zap.Error(err))
		return nil, err
	}

	resp := model.LoginResp{
		Token: token,
		User: model.GetUserInfoResp{
			Uid:      user.Id,
			Name:     user.Name,
			Address:  user.Address,
			Phone:    user.Phone,
			Identity: user.Identity,
		},
	}

	log.Infow("success!")
	return resp, nil
}

// 发送验证码
func (s *Service) SendVerificationCode(ctx context.Context, req model.SendVerificationCodeReq) (interface{}, error) {
	log := logging.For(ctx, "func", "SendVerificationCode",
		zap.Any("req", req),
	)

	//todo 生成验证码用随机数,调用真实接口发送短信验证码
	code := "1234"
	err := s.dao.SetVerificationCode(ctx, req.Phone, code)
	if err != nil {
		log.Errorw("s.dao.SetVerificationCode err", zap.Error(err))
		return nil, err
	}

	log.Infow("success!")
	return nil, nil
}

// 刷新token时间
func (s *Service) RefreshToken(ctx context.Context, token string) error {
	log := logging.For(ctx, "func", "RefreshToken",
		zap.String("token", token),
	)
	err := s.dao.RefreshTokenExpire(ctx, token)
	if err != nil {
		log.Errorw("s.dao.SetVerificationCode err", zap.Error(err))
		return err
	}

	log.Infow("success!")
	return nil
}
