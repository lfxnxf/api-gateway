package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

// 新增班次
func (d *Dao) InsertShifts(ctx context.Context, tx *gorm.DB, m model.ShiftsModel) (model.ShiftsModel, error) {
	log := logging.For(ctx, "func", "InsertShifts",
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
		return m, err
	}
	log.Infow("success")
	return m, nil
}

// 删除班次
func (d *Dao) DeleteShifts(ctx context.Context, tx *gorm.DB, id int64) error {
	log := logging.For(ctx, "func", "DeleteShifts",
		zap.Int64("id", id),
	)
	if tx == nil {
		tx = d.db.Master(ctx).DB
	}
	err := tx.Table(model.ShiftsTableName).Where("id = ?", id).Update(map[string]interface{}{"status": model.ShiftsStatusDeleted}).Error
	if err != nil {
		log.Errorw("Delete",
			zap.String("err", err.Error()),
		)
		return err
	}
	log.Infow("success")
	return nil
}

