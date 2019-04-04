package config

import (
	log "airman.com/airms/pkg/mslog"
	"github.com/BurntSushi/toml"
)

var _Config Config

type Config struct {
	Logger   *Logger   `toml:"log"`
	Service  *Service  `toml:"service"`
	Register *Register `toml:"register"`
	Trace    *Trace    `toml:"trace"`
	Biz      *Biz      `toml:"biz"`
}

type Logger struct {
	Level string `toml:"level"`
	Size  int64  `toml:"size"`
	Name  string `toml:"name"`
	Path  string `toml:"path"`
}

type Service struct {
	Name          string `toml:"name"`
	Address       string `toml:"address"`
	MetricFlag    bool   `toml:"isMetric"`
	MetricAddress string `toml:"metricAddress"`
}

type Biz struct {
	MysqlFile string `toml:"mysql"`
	RedisFile string `toml:"redis"`
	Root      string `toml:"root"`
}

type Register struct {
	Addresses []string `toml:"addresses"`
	TTL       int      `toml:"ttl"`
	Timeout   int      `toml:"timeout"`
}

type Trace struct {
	Url string `toml:"url"`
}

func GetLogger() *Logger {
	return _Config.Logger
}

func GetRegister() *Register {
	return _Config.Register
}

func GetTrace() *Trace {
	return _Config.Trace
}

func GetBiz() *Biz {
	return _Config.Biz
}

func GetService() *Service {
	return _Config.Service
}

func Setup(fileName string) error {
	if _, err := toml.DecodeFile(fileName, &_Config); err != nil {
		return err

	}
	log.Infof("config: %#+v", _Config)
	return nil
}
