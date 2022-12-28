package global

import (
	"github.com/geekr-dev/go-blog-app/pkg/config"
	"github.com/geekr-dev/go-blog-app/pkg/logger"
)

var (
	ServerConfig   *config.ServerConfig
	AppConfig      *config.AppConfig
	DatabaseConfig *config.DatabaseConfig

	Logger *logger.Logger
)
