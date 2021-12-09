package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

// 新增班次
func (d *Dao) InsertShiftsSites(ctx context.Context, tx *gorm.DB, m model.ShiftsSitesModel) error {
	log := logging.For(ctx, "func", "InsertShiftsSites",
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

// 删除全部班次站点
func (d *Dao) DeleteShiftsSites(ctx context.Context, tx *gorm.DB, shiftId int64) error {
	log := logging.For(ctx, "func", "DeleteShiftsSites",
		zap.Int64("shift_id", shiftId),
	)
	if tx == nil {
		tx = d.db.Master(ctx).DB
	}
	err := tx.Table(model.ShiftsSitesTableName).Where("shift_id = ?", shiftId).Update(map[string]interface{}{"status": model.ShiftsStatusDeleted}).Error
	if err != nil {
		log.Errorw("Delete",
			zap.String("err", err.Error()),
		)
		return err
	}
	log.Infow("success")
	return nil
}

// 删除某个次站点
func (d *Dao) DeleteShiftsSitesById(ctx context.Context, tx *gorm.DB, id int64) error {
	log := logging.For(ctx, "func", "DeleteShiftsSitesById",
		zap.Int64("id", id),
	)
	if tx == nil {
		tx = d.db.Master(ctx).DB
	}
	err := tx.Table(model.ShiftsSitesTableName).Where("id = ?", id).Update(map[string]interface{}{"status": model.ShiftsStatusDeleted}).Error
	if err != nil {
		log.Errorw("Delete",
			zap.String("err", err.Error()),
		)
		return err
	}
	log.Infow("success")
	return nil
}


