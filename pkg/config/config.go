package config

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var sections = make(map[string]interface{})

type ServerConfig struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppConfig struct {
	Name                 string
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DatabaseConfig struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTConfig struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailConfig struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
}

type Config struct {
	*viper.Viper
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("configs/")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	c := &Config{v}
	c.WatchConfigChange()
	return c, nil
}

func (c *Config) ReadSection(k string, v interface{}) error {
	err := c.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (c *Config) ReloadAllSection() error {
	for k, v := range sections {
		err := c.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) WatchConfigChange() {
	go func() {
		c.WatchConfig()
		c.OnConfigChange(func(e fsnotify.Event) {
			_ = c.ReloadAllSection()
		})
	}()
}
