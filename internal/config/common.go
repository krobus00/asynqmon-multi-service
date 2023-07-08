package config

type ServiceConfig struct {
	Name      string `json:"name"`
	RedisHost string `json:"redis_host"`
}
