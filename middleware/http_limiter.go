package middleware

import (
	"net/http"
	"strings"

	"github.com/k3ndzy/go-rate-limiter/limiter"
)

// RateLimitMiddleware tích hợp HTTP headers chuẩn RFC 6585
func RateLimitMiddleware(rl *limiter.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP := r.RemoteAddr
			if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
				clientIP = strings.Split(forwarded, ",")[0]
			}

			if !rl.Allow(clientIP) {
				w.Header().Set("Retry-After", "1")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = w.Write([]byte(`{"error":"quá nhiều yêu cầu, vui lòng thử lại sau","status":429}`))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
