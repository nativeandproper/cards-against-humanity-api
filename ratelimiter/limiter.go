package ratelimiter

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"math"
	"time"
)

type Limiter struct {
	client *redis.Client
	*LimiterOptions
}

type LimiterOptions struct {
	SlidingWindow time.Duration
	Limit         int64
	Namespace     string
}

// New returns an instance of the rate limiter client
func New(redisClient *redis.Client, options *LimiterOptions) *Limiter {
	return &Limiter{
		client:         redisClient,
		LimiterOptions: options,
	}
}

// Enforce enforces rate limit
func (l *Limiter) Enforce(key string) (int64, error) {
	if l.Limit == 0 {
		return 0, errors.New("no limit set")
	}

	prefixedKey := l.prefixKey(key)

	now := time.Now()
	prevWindow := now.Add(-l.SlidingWindow)
	member := redis.Z{
		Score:  float64(now.Unix()),
		Member: float64(now.Unix()),
	}

	pipe := l.client.TxPipeline()
	pipe.ZRangeByScoreWithScores(prefixedKey, redis.ZRangeBy{
		Min: fmt.Sprintf("%f", float64(0)),
		Max: fmt.Sprintf("%f", float64(prevWindow.Unix())),
	})
	// TODO: check count before incrementing in LUA script
	pipe.ZAdd(prefixedKey, member)
	pipe.ZCount(prefixedKey, fmt.Sprintf("%f", float64(0)), fmt.Sprintf("%f", float64(now.Unix())))
	pipe.Expire(prefixedKey, l.SlidingWindow)

	cmds, err := pipe.Exec()
	if err != nil {
		return 0, err
	}

	var remaining int64
	if len(cmds) >= 2 {
		reqCount := cmds[2].(*redis.IntCmd).Val()
		remaining = int64(math.Max(float64(0), float64(l.Limit-reqCount)))
	}

	return remaining, nil
}

func (l *Limiter) prefixKey(key string) string {
	return l.Namespace + key
}
