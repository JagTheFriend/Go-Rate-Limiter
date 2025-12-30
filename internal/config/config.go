package config

import "time"

type Config struct {
	RedisAddr string
	Capacity  int
	Refill    int
	Interval  time.Duration
}

func Default() Config {
	return Config{
		RedisAddr: "localhost:6379",
		Capacity:  10,
		Refill:    10,
		Interval:  time.Minute,
	}
}
