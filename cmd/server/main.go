package main

import (
	"log"
	"net/http"

	"go-rate-limiter/internal/api"
	"go-rate-limiter/internal/config"
	"go-rate-limiter/internal/limiter"
	"go-rate-limiter/internal/middleware"
	"go-rate-limiter/internal/storage"
)

func main() {
	cfg := config.Default()

	redis := storage.New(cfg.RedisAddr)

	tb := &limiter.TokenBucket{
		Redis:    redis,
		Capacity: cfg.Capacity,
		Refill:   cfg.Refill,
		Interval: cfg.Interval,
	}

	mux := http.NewServeMux()
	mux.Handle("/allow", api.AllowHandler(tb))

	handler := middleware.RateLimit(tb)(mux)

	log.Println("Rate Limiter running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
