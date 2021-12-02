package dao

import (
	"context"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

// 获取用户信息
func (d *Dao) GetUserInfo(ctx context.Context, uid int64) (model.UsersModel, error) {
	log := logging.For(ctx, "func", "GetUserInfo",
		zap.Int64("uid", uid),
	)
	var users []model.UsersModel
	err := d.db.Slave(ctx).Table(model.UsersTableName).Where("id = ? and status = ? ", uid, model.UserStatusNormal).Scan(&users).Error
	if err != nil {
		log.Errorw("Get",
			zap.String("err", err.Error()),
		)
		return model.UsersModel{}, err
	}

	if len(users) <= 0 {
		return model.UsersModel{}, nil
	}

	log.Infow("success", zap.Any("users", users))
	return users[0], nil
}

// 通过手机号获取用户信息
func (d *Dao) GetUserInfoByPhone(ctx context.Context, phone int64) (model.UsersModel, error) {
	log := logging.For(ctx, "func", "GetUserInfoByPhone",
		zap.Int64("phone", phone),
	)
	var users []model.UsersModel
	err := d.db.Slave(ctx).Table(model.UsersTableName).Where("phone = ? and status = ? ", phone, model.UserStatusNormal).Scan(&users).Error
	if err != nil {
		log.Errorw("Get",
			zap.String("err", err.Error()),
		)
		return model.UsersModel{}, err
	}

	if len(users) <= 0 {
		return model.UsersModel{}, nil
	}

	log.Infow("success", zap.Any("users", users))
	return users[0], nil
}

// 获取司机信息
func (d *Dao) GetDrivers(ctx context.Context, uid int64) ([]model.UsersModel, error) {
	log := logging.For(ctx, "func", "GetDrivers",
		zap.Int64("uid", uid),
	)
	var resp []model.UsersModel
	err := d.db.Slave(ctx).Table(model.UsersTableName).Where("((id = ? and identity = ?) or (boss_id = ? and identity = ?)) and status = ? ", uid, model.IdentityBoss, uid, model.IdentityDriver, model.UserStatusNormal).Scan(&resp).Error
	if err != nil {
		log.Errorw("Get",
			zap.String("err", err.Error()),
		)
		return resp, err
	}
	log.Infow("success")
	return resp, nil
}

// 新增用户
func (d *Dao) InsertUsers(ctx context.Context, m model.UsersModel) (model.UsersModel, error) {
	log := logging.For(ctx, "func", "InsertUsers",
		zap.Any("model", m),
	)
	err := d.db.Master(ctx).Create(&m).Error
	if err != nil {
		log.Errorw("Create",
			zap.String("err", err.Error()),
		)
		return m, err
	}
	log.Infow("success")
	return m, nil
}

func (d *Dao) UpdateUsersByMap(ctx context.Context, id int64, data map[string]interface{}) error {
	log := logging.For(ctx, "func", "InsertUsers",
		zap.Int64("uid", id),
		zap.Any("data", data),
	)
	err := d.db.Master(ctx).Table(model.UsersTableName).Where("id = ?", id).Update(data).Error
	if err != nil {
		log.Errorw("Update",
			zap.String("err", err.Error()),
		)
		return err
	}
	log.Infow("success")
	return nil
}