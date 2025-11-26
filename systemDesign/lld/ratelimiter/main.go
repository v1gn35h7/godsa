package ratelimiter

import (
	"fmt"

	"github.com/vicky1115050/godsa/systemDesign/lld/ratelimiter/application"
)

func main() {
	fmt.Println("Rate Limiter")

	// start schudler

	// start api server
	rateLimiterFactory := &application.RateLimiterFactory{}
	ratelimiter := rateLimiterFactory.CreateRateLimiter("tokenBucket")

	server := application.NewRLServer(ratelimiter)

	for _, user := range []string{"vicky", "venkat", "laxman", "sekar", "unknown"} {
		if server.HandleRequest(user) {
			fmt.Printf("Request from %s is allowed\n", user)
		} else {
			fmt.Printf("Request from %s is denied\n", user)
		}
	}

}
