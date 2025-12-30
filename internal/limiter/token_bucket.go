package limiter

import (
	"context"
	"strconv"
	"time"

	"go-rate-limiter/internal/storage"
)

type TokenBucket struct {
	Redis    *storage.Redis
	Capacity int
	Refill   int
	Interval time.Duration
}

func (tb *TokenBucket) Allow(ctx context.Context, key string) (bool, int, error) {
	now := time.Now().Unix()

	tokenKey := "tokens:" + key
	timeKey := "last:" + key

	lastTimeStr, _ := tb.Redis.Get(ctx, timeKey)
	lastTime, _ := strconv.ParseInt(lastTimeStr, 10, 64)

	elapsed := now - lastTime
	refillTokens := int(elapsed/int64(tb.Interval.Seconds())) * tb.Refill

	tokenStr, _ := tb.Redis.Get(ctx, tokenKey)
	tokens, _ := strconv.Atoi(tokenStr)

	if tokens == 0 {
		tokens = tb.Capacity
	}

	tokens = min(tb.Capacity, tokens+refillTokens)

	if tokens <= 0 {
		return false, 0, nil
	}

	tokens--

	err := tb.Redis.Set(ctx, tokenKey, tokens, int64(tb.Interval.Seconds()))
	if err != nil {
		return false, 0, err
	}

	err = tb.Redis.Set(ctx, timeKey, now, int64(tb.Interval.Seconds()))
	if err != nil {
		return false, 0, err
	}

	return true, tokens, nil
}
