package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
	"time"
)

const (
	tokenKey       = "school:token:string"   // token
	tokenKeyExpire = 7 * 86400 * time.Second // token过期时间

	verificationCodeKey       = "school:verification:code:string" // 验证码
	verificationCodeKeyExpire = 300 * time.Second                 // 验证码有效期，5分钟
)

func (d *Dao) tokenKey(token string) string {
	return fmt.Sprintf("%s:%s", tokenKey, token)
}

func (d *Dao) verificationCodeKey(phone int64) string {
	return fmt.Sprintf("%s:%d", verificationCodeKey, phone)
}

// 设置token
func (d *Dao) SetToken(ctx context.Context, token string, user model.UsersModel) error {
	log := logging.For(ctx, "func", "SetToken",
		zap.String("token", token),
		zap.Any("m", user),
	)

	bytesVal, _ := json.Marshal(user)
	key := d.tokenKey(token)
	_, err := d.cache.Set(ctx, key, string(bytesVal))

	if err != nil {
		log.Errorw("d.cache.Set error", zap.Error(err))
		return err
	}

	_ = d.cache.Expire(ctx, key, tokenKeyExpire)

	log.Infow("success!")
	return nil
}

// 刷新token过期时间
func (d *Dao) RefreshTokenExpire(ctx context.Context, token string) error {
	log := logging.For(ctx, "func", "RefreshTokenExpire",
		zap.String("token", token),
	)
	key := d.tokenKey(token)
	err := d.cache.Expire(ctx, key, tokenKeyExpire)
	if err != nil {
		log.Errorw("d.cache.Set error", zap.Error(err))
		return err
	}
	log.Infow("success!")
	return nil
}

// 通过token获取用户信息
func (d *Dao) GetUserByToken(ctx context.Context, token string) (model.UsersModel, error) {
	log := logging.For(ctx, "func", "GetUserByToken",
		zap.String("token", token),
	)
	key := d.tokenKey(token)
	var resp model.UsersModel
	bytesVal, err := d.cache.Get(ctx, key)
	if err != nil {
		log.Errorw("d.cache.Set error", zap.Error(err))
		return resp, err
	}
	_ = json.Unmarshal(bytesVal, &resp)

	log.Infow("success!", zap.Any("resp", resp))
	return resp, nil
}

// 设置验证码
func (d *Dao) SetVerificationCode(ctx context.Context, phone int64, code string) error {
	log := logging.For(ctx, "func", "SetToken",
		zap.Int64("phone", phone),
		zap.String("code", code),
	)
	key := d.verificationCodeKey(phone)
	_, err := d.cache.Set(ctx, key, code)

	if err != nil {
		log.Errorw("d.cache.Set error", zap.Error(err))
		return err
	}

	_ = d.cache.Expire(ctx, key, verificationCodeKeyExpire)

	log.Infow("success!")
	return nil
}

// 获取验证码
func (d *Dao) GetVerificationCode(ctx context.Context, phone int64) (string, error) {
	log := logging.For(ctx, "func", "SetToken",
		zap.Int64("phone", phone),
	)
	key := d.verificationCodeKey(phone)
	bytesVal, err := d.cache.Get(ctx, key)

	if err != nil {
		log.Errorw("d.cache.Get error", zap.Error(err))
		return "", err
	}

	log.Infow("success!")
	return string(bytesVal), nil
}
