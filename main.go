package main

import (
	"log"
	"net/http"
	"time"

	"github.com/geekr-dev/go-blog-app/global"
	"github.com/geekr-dev/go-blog-app/internal/model"
	"github.com/geekr-dev/go-blog-app/internal/routers"
	"github.com/geekr-dev/go-blog-app/pkg/config"
	"github.com/geekr-dev/go-blog-app/pkg/logger"
	"github.com/geekr-dev/go-blog-app/pkg/tracer"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// 初始化全局配置
	err := setupConfig()
	if err != nil {
		log.Fatalf("init.setupConfig err: %v", err)
	}
	// 初始化全局日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	// global.Logger.Infof("%s inited", "blog-app")
	// 初始化链路追踪
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	// 初始化数据库连接
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

// @title Blog App
// @version 1.0
// @description 基于 Go 构建简单博客系统
// @termsOfService https://github.com/geekr-dev/go-blog-app
func main() {
	gin.SetMode(global.ServerConfig.RunMode)
	router := routers.NewRouter()

	s := http.Server{
		Addr:           ":" + global.ServerConfig.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerConfig.ReadTimeout,
		WriteTimeout:   global.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func setupConfig() error {
	config, err := config.NewConfig()
	if err != nil {
		return err
	}
	err = config.ReadSection("Server", &global.ServerConfig)
	if err != nil {
		return err
	}
	err = config.ReadSection("App", &global.AppConfig)
	if err != nil {
		return err
	}
	err = config.ReadSection("Database", &global.DatabaseConfig)
	if err != nil {
		return err
	}
	err = config.ReadSection("JWT", &global.JWTConfig)
	if err != nil {
		return err
	}
	err = config.ReadSection("Email", &global.EmailConfig)
	if err != nil {
		return err
	}
	global.ServerConfig.ReadTimeout *= time.Second
	global.ServerConfig.WriteTimeout *= time.Second
	global.JWTConfig.Expire *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseConfig)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  global.AppConfig.LogSavePath + "/" + global.AppConfig.LogFileName + global.AppConfig.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		},
		"",
		log.LstdFlags,
	).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(global.AppConfig.Name, "localhost:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
