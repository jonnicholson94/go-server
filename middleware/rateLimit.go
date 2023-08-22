package middleware

import (
	"net/http"
	"time"
)

type RateLimiter struct {
	limiter <-chan time.Time
}

func NewRateLimiter(rate int) *RateLimiter {
	timeInterval := time.Second / time.Duration(rate)
	return &RateLimiter{
		limiter: time.Tick(timeInterval),
	}
}

func (r1 *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r1.limiter
		next.ServeHTTP(w, r)
	})
}
