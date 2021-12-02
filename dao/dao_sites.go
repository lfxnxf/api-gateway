package dao

import (
	"context"
	"fmt"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

func (d *Dao) GetAllSites(ctx context.Context, where map[string]interface{}) ([]model.SitesModel, error) {
	log := logging.For(ctx, "func", "GetAllSites",
		zap.Any("where", where),
	)
	var resp []model.SitesModel

	query := d.db.Slave(ctx).Table(model.SitesTableName).Where("1 = 1")

	for k, v := range where {
		if k == "name" {
			query = query.Where(fmt.Sprintf("%s like ?", k), fmt.Sprintf("%%%s%%", v))
			continue
		}
		query = query.Where(fmt.Sprintf("%s = ?", k), v)
	}

	err := d.db.Slave(ctx).Table(model.SitesTableName).
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
