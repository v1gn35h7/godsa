package application

import (
	"time"

	"github.com/vicky1115050/godsa/systemDesign/lld/ratelimiter/entites"
)

type tokenBucketStrategy struct {
	buckets map[string]*entites.TokenBucket

	// global bucket
	globalBucket *entites.TokenBucket
}

func NewTokenBucketStrategy(globalCapacity, userCapacity int, users []string) RateLimitStrategy {
	s := &tokenBucketStrategy{
		buckets:      make(map[string]*entites.TokenBucket),
		globalBucket: &entites.TokenBucket{Capacitiy: globalCapacity},
	}

	for _, v := range users {
		s.buckets[v] = &entites.TokenBucket{
			Key:       v,
			Capacitiy: userCapacity,
		}
	}

	return s
}

func (t *tokenBucketStrategy) GiveAccess(key string) bool {
	if t.GetGlobalToken() && t.GetUserToken(key) {
		return true
	}
	return false
}

func (s *tokenBucketStrategy) GetGlobalToken() bool {
	if s.globalBucket.Capacitiy > 0 {
		s.globalBucket.MU.Lock()
		defer s.globalBucket.MU.Unlock()
		s.globalBucket.Capacitiy--
		return true
	}

	return false
}

func (s *tokenBucketStrategy) GetUserToken(key string) bool {
	if bucket, ok := s.buckets[key]; ok {
		if bucket.Capacitiy <= 0 {
			return false
		}
		bucket.MU.Lock()
		defer bucket.MU.Unlock()
		bucket.Capacitiy--
		return true
	}
	return false
}

func (s *tokenBucketStrategy) refillGlobalToken(tokens int) {
	s.globalBucket.MU.Lock()
	defer s.globalBucket.MU.Unlock()
	s.globalBucket.Capacitiy += tokens
}

func (s *tokenBucketStrategy) refillUserToken(key string, tokens int) {
	if bucket, ok := s.buckets[key]; ok {
		bucket.MU.Lock()
		defer bucket.MU.Unlock()
		bucket.Capacitiy += tokens
	}
}

func (s *tokenBucketStrategy) SchudeuleRefillGlobalToken(interval int) {
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.refillGlobalToken(10)
				for _, userBucket := range s.buckets {
					s.refillUserToken(userBucket.Key, 5)
				}
			}
		}
	}()
}
