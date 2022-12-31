package model

import (
	"fmt"
	"time"

	"github.com/geekr-dev/go-blog-app/global"
	"github.com/geekr-dev/go-blog-app/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormopentracing "gorm.io/plugin/opentracing"
)

const (
	STATE_CLOSE = iota
	STATE_OPEN
)

type Model struct {
	ID        uint32         `gorm:"primary_key" json:"id,omitempty"`
	CreatedBy string         `json:"created_by,omitempty"`
	UpdatedBy string         `json:"updated_by,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func NewDBEngine(dbConfig *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
				dbConfig.UserName,
				dbConfig.Password,
				dbConfig.Host,
				dbConfig.DBName,
				dbConfig.Charset,
				dbConfig.ParseTime,
			),
		}),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	// 设置日志模式
	if global.ServerConfig.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}
	// 设置闲置/最大连接数
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	// 设置链路追踪
	db.Use(gormopentracing.New())
	return db, nil
}
