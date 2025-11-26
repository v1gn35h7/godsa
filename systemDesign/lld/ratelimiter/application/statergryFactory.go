package application

type RateLimiterFactory struct {
}

func (f *RateLimiterFactory) CreateRateLimiter(strategyType string) RateLimitStrategy {
	if strategyType == "tokenBucket" {
		return NewTokenBucketStrategy(100, 10, []string{"user1", "user2", "user3"})
	}

	return nil
}
