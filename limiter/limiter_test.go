package limiter

import (
	"testing"
	"time"
)

func TestBucket_AllowAndRefill(t *testing.T) {
	// Dung lượng 3 token, nạp 10 token/giây
	b := NewBucket(3, 10)

	// Tiêu thụ 3 token đầu tiên
	if !b.Allow() || !b.Allow() || !b.Allow() {
		t.Fatal("phải cho phép 3 request đầu tiên")
	}

	// Request thứ 4 phải bị chặn ngay lập tức
	if b.Allow() {
		t.Fatal("request thứ 4 phải bị từ chối")
	}

	// Đợi 150ms (nạp lại ít nhất 1.5 token)
	time.Sleep(150 * time.Millisecond)

	if !b.Allow() {
		t.Fatal("sau khi chờ nạp token, request phải được chấp thuận")
	}
}

func TestRateLimiter_MultipleClients(t *testing.T) {
	rl := New(2, 1)

	// Client 1 dùng hết quota
	if !rl.Allow("client-A") || !rl.Allow("client-A") {
		t.Fatal("client A phải được phép dùng 2 request")
	}
	if rl.Allow("client-A") {
		t.Fatal("client A phải bị chặn ở request thứ 3")
	}

	// Client 2 phải hoàn toàn độc lập và không bị ảnh hưởng
	if !rl.Allow("client-B") {
		t.Fatal("client B phải được phép gửi request")
	}
}
