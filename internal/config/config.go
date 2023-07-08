package config

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/krobus00/asynqmon-multi-service/internal/util"
	"github.com/spf13/viper"
)

var (
	serviceName    = ""
	serviceVersion = ""
)

func ServiceName() string {
	return serviceName
}

func ServiceVersion() string {
	return serviceVersion
}

func Env() string {
	return viper.GetString("env")
}

func LogLevel() string {
	return viper.GetString("log_level")
}

func PortHTTP() string {
	return viper.GetString("ports.http")
}

func GracefulShutdownTimeOut() time.Duration {
	cfg := viper.GetString("graceful_shutdown_timeout")
	return parseDuration(cfg, DefaultGracefulShutdownTimeOut)
}

func MonitoringPath() string {
	return viper.GetString("monitoring_path")
}

func Services() []ServiceConfig {
	services := make([]ServiceConfig, 0)
	jsonData, err := json.Marshal(viper.Get("services"))
	util.ContinueOrFatal(err)
	err = json.Unmarshal(jsonData, &services)
	util.ContinueOrFatal(err)
	return services
}

func parseDuration(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}
	return dur
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config not found")
		}
		return err
	}
	return nil
}
