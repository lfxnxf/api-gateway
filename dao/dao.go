package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/lfxnxf/frame/logic/inits/proxy"

	"github.com/lfxnxf/school/api-gateway/conf"
)

const (
	RedisClient = "school.api.redis"
	DbClient = "school.api.db"
)

// Dao represents data access object
type Dao struct {
	c     *conf.Config
	db    *proxy.SQL
	cache *proxy.Redis
}

func New(c *conf.Config) *Dao {
	return &Dao{
		c:     c,
		db:    proxy.InitSQL(DbClient),
		cache: proxy.InitRedis(RedisClient),
	}
}

// Ping check db resource status
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}

// 开启事务
func (d *Dao) StartTransaction(ctx context.Context) *gorm.DB {
	return d.db.Master(ctx).Begin()
}
