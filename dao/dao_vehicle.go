package dao

import (
	"context"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

func (d *Dao) GetVehicleInfoByLicensePlate(ctx context.Context, licensePlate string) (model.VehicleModel, error) {
	log := logging.For(ctx, "func", "GetVehicleInfoByLicensePlate",
		zap.String("license_plate", licensePlate),
	)
	var vehicles []model.VehicleModel
	err := d.db.Slave(ctx).Table(model.VehicleTableName).Where("license_plate = ? and status = ? ", licensePlate, model.VehicleNormal).Scan(&vehicles).Error
	if err != nil {
		log.Errorw("Get",
			zap.String("err", err.Error()),
		)
		return model.VehicleModel{}, err
	}

	if len(vehicles) <= 0 {
		return model.VehicleModel{}, nil
	}

	log.Infow("success", zap.Any("users", vehicles))
	return vehicles[0], nil
}

func (d *Dao) GetVehicleInfoById(ctx context.Context, id int64) (model.VehicleModel, error) {
	log := logging.For(ctx, "func", "GetUserInfo",
		zap.Int64("id", id),
	)
	var vehicles []model.VehicleModel
	err := d.db.Slave(ctx).Table(model.VehicleTableName).Where("id = ? and status = ? ", id, model.VehicleNormal).Scan(&vehicles).Error
	if err != nil {
		log.Errorw("Get",
			zap.String("err", err.Error()),
		)
		return model.VehicleModel{}, err
	}

	if len(vehicles) <= 0 {
		return model.VehicleModel{}, nil
	}

	log.Infow("success", zap.Any("users", vehicles))
	return vehicles[0], nil
}

// 新增车辆
func (d *Dao) InsertVehicle(ctx context.Context, m model.VehicleModel) error {
	log := logging.For(ctx, "func", "InsertVehicle",
		zap.Any("model", m),
	)
	err := d.db.Master(ctx).Create(&m).Error
	if err != nil {
		log.Errorw("Create",
			zap.String("err", err.Error()),
		)
		return err
	}
	log.Infow("success")
	return nil
}

// 修改车辆
func (d *Dao) UpdateVehicle(ctx context.Context, m model.VehicleModel) error {
	log := logging.For(ctx, "func", "UpdateVehicle",
		zap.Any("m", m),
	)
	err := d.db.Master(ctx).Table(model.VehicleTableName).Table(model.VehicleTableName).Where("id = ?", m.Id).Update(&m).Error
	if err != nil {
		log.Errorw("Update",
			zap.String("err", err.Error()),
		)
		return err
	}
	log.Infow("success")
	return nil
}

// 通过老板获取车辆
func (d *Dao) GetAllVehicleByBoss(ctx context.Context, bossUid int64) ([]model.VehicleModel, error) {
	log := logging.For(ctx, "func", "GetAllVehicleByBoss",
		zap.Int64("boss_uid", bossUid),
	)
	var resp []model.VehicleModel
	err := d.db.Slave(ctx).Table(model.VehicleTableName).Table(model.VehicleTableName).
		Where("boss_id = ? and status = ?", bossUid, model.VehicleNormal).
		Scan(&resp).Error
	if err != nil {
		log.Errorw("Get",
			zap.String("err", err.Error()),
		)
		return resp, err
	}
	log.Infow("success", zap.Any("resp", resp))
	return resp, nil
}

func (d *Dao) GetAllVehicleInfo(ctx context.Context) ([]model.VehicleInfoModel, error) {
	log := logging.For(ctx, "func", "GetAllVehicleInfo")
	var resp []model.VehicleInfoModel
	err := d.db.Slave(ctx).Table(model.VehicleInfoTableName).
		Scan(&resp).Error
	if err != nil {
		log.Errorw("Get",
			zap.String("err", err.Error()),
		)
		return resp, err
	}
	log.Infow("success", zap.Any("resp", resp))
	return resp, nil
}