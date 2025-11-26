package application

type RateLimitStrategy interface {
	GiveAccess(key string) bool
}
