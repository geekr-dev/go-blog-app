package main

import (
	"log"
	"net/http"
	"time"

	"github.com/geekr-dev/go-blog-app/global"
	"github.com/geekr-dev/go-blog-app/internal/routers"
	"github.com/geekr-dev/go-blog-app/pkg/config"
	"github.com/gin-gonic/gin"
)

func init() {
	err := setupConfig()
	if err != nil {
		log.Fatalf("init.setupConfig err: %v", err)
	}
}

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
	global.ServerConfig.ReadTimeout *= time.Second
	global.ServerConfig.WriteTimeout *= time.Second

	return nil
}
