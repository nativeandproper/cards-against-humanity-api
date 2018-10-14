package ratelimiter

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
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

// Enforce enforces the rate limit and remaining requests
func (l *Limiter) Enforce(key string) (int64, error) {
	if l.Limit == 0 {
		return 0, errors.New("no limit set")
	}

	pKey := l.prefixKey(key)
	now := time.Now().Unix()

	minScore := stringifyInt64AsFloat(0)
	prevMaxScore := stringifyInt64AsFloat(time.Now().Add(-l.SlidingWindow).Unix())
	currMaxScore := stringifyInt64AsFloat(now)

	member := redis.Z{
		Score:  float64(now),
		Member: float64(now),
	}

	rangeBy := redis.ZRangeBy{
		Min: minScore,
		Max: prevMaxScore,
	}

	// Execute transaction
	pipe := l.client.TxPipeline()
	pipe.ZRangeByScoreWithScores(pKey, rangeBy)
	pipe.ZAdd(pKey, member)
	pipe.ZCount(pKey, minScore, currMaxScore)
	pipe.Expire(pKey, l.SlidingWindow)

	cmds, err := pipe.Exec()
	if err != nil {
		return 0, err
	}

	var remReq int64
	if len(cmds) >= 2 {
		zAddResult := cmds[2]
		totalReq := zAddResult.(*redis.IntCmd).Val()
		remReq = int64(math.Max(float64(0), float64(l.Limit-totalReq)))
	}

	return remReq, nil
}

func stringifyInt64AsFloat(num int64) string {
	return strconv.FormatFloat(float64(num), 'f', 0, 64)
}

func (l *Limiter) prefixKey(key string) string {
	return l.Namespace + key
}
