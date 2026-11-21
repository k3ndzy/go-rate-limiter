package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/k3ndzy/go-rate-limiter/limiter"
	"github.com/k3ndzy/go-rate-limiter/middleware"
)

func main() {
	// Cho phép tối đa 5 request burst, nạp lại 2 request/giây
	rl := limiter.New(5, 2)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"Xin chào từ server! Request của bạn đã được thông qua.","status":"success"}`)
	})

	handler := middleware.RateLimitMiddleware(rl)(mux)

	addr := ":8082"
	fmt.Printf("🛡️ Rate Limiter API Server đang lắng nghe tại http://localhost%s\n", addr)
	fmt.Println("👉 Thử spam F5 liên tục tại http://localhost:8082/api/hello để xem mã lỗi 429!")

	log.Fatal(http.ListenAndServe(addr, handler))
}
