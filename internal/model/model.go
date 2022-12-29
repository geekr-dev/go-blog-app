package model

import (
	"fmt"
	"time"

	"github.com/geekr-dev/go-blog-app/global"
	"github.com/geekr-dev/go-blog-app/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID        uint32         `gorm:"primary_key" json:"id"`
	CreatedBy string         `json:"created_by"`
	UpdatedBy string         `json:"updated_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
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
	if global.ServerConfig.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	return db, nil
}
