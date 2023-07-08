package config

import "time"

const (
	DefaultGracefulShutdownTimeOut = 30 * time.Second

	DefaultRedisDialTimeout  = 5 * time.Second
	DefaultRedisWriteTimeout = 2 * time.Second
	DefaultRedisReadTimeout  = 2 * time.Second
)
