package limiter

import (
	"sync"
	"time"
)

// Bucket đại diện cho Token Bucket của một client/IP
type Bucket struct {
	capacity   float64
	tokens     float64
	refillRate float64 // Số token được nạp mỗi giây
	lastRefill time.Time
	mu         sync.Mutex
}

func NewBucket(capacity, refillRate float64) *Bucket {
	return &Bucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow kiểm tra và tiêu thụ 1 token nếu còn đủ
func (b *Bucket) Allow() bool {
	return b.AllowN(1)
}

func (b *Bucket) AllowN(n float64) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.lastRefill = now

	// Nạp thêm token dựa theo thời gian trôi qua
	b.tokens += elapsed * b.refillRate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}

	if b.tokens >= n {
		b.tokens -= n
		return true
	}

	return false
}

// RateLimiter quản lý bucket cho từng client IP
type RateLimiter struct {
	capacity   float64
	refillRate float64
	buckets    map[string]*Bucket
	mu         sync.RWMutex
}

func New(capacity, refillRate float64) *RateLimiter {
	return &RateLimiter{
		capacity:   capacity,
		refillRate: refillRate,
		buckets:    make(map[string]*Bucket),
	}
}

func (rl *RateLimiter) Allow(clientKey string) bool {
	rl.mu.RLock()
	b, exists := rl.buckets[clientKey]
	rl.mu.RUnlock()

	if !exists {
		rl.mu.Lock()
		b, exists = rl.buckets[clientKey]
		if !exists {
			b = NewBucket(rl.capacity, rl.refillRate)
			rl.buckets[clientKey] = b
		}
		rl.mu.Unlock()
	}

	return b.Allow()
}
