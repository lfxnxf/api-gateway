package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

// 车辆新增站点
func (d *Dao) InsertVehicleSites(ctx context.Context, tx *gorm.DB, m []model.VehicleSitesModel) error {
	log := logging.For(ctx, "func", "InsertVehicleSites",
		zap.Any("model", m),
	)
	if tx == nil {
		tx = d.db.Master(ctx).DB
	}
	err := tx.Create(&m).Error
	if err != nil {
		log.Errorw("Create",
			zap.String("err", err.Error()),
		)
		return err
	}
	log.Infow("success")
	return nil
}

// 删除车辆全部站点
func (d *Dao) DeleteVehicleSites(ctx context.Context, tx *gorm.DB, vehicleId int64) error {
	log := logging.For(ctx, "func", "DeleteVehicleSites",
		zap.Any("vehicle_id", vehicleId),
	)
	if tx == nil {
		tx = d.db.Master(ctx).DB
	}
	err := tx.Where("vehicle_id = ?", vehicleId).Update(map[string]interface{}{
		"status": model.VehicleSitesStatusDeleted,
	}).Error
	if err != nil {
		log.Errorw("Delete",
			zap.String("err", err.Error()),
		)
		return err
	}
	log.Infow("success")
	return nil
}

func (d *Dao) GetVehicleSites(ctx context.Context, vehicleId int64) ([]model.VehicleSitesModel, error) {
	log := logging.For(ctx, "func", "GetVehicleSites",
		zap.Int64("vehicle_id", vehicleId),
	)
	var resp []model.VehicleSitesModel
	err := d.db.Slave(ctx).Table(model.VehicleSitesTableName).
		Where("vehicle_id = ? and status = ?", vehicleId, model.VehicleSitesStatusNormal).Scan(&resp).Error
	if err != nil {
		log.Errorw("Get",
			zap.String("err", err.Error()),
		)
		return resp, err
	}
	log.Infow("success")
	return resp, nil
}

