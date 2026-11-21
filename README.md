# 🛡️ Go Rate Limiter: Production Token Bucket Middleware

Thư viện và middleware giới hạn tần suất yêu cầu (Rate Limiter) chuẩn thuật toán **Token Bucket** bằng Go, bảo vệ hệ thống khỏi các cuộc tấn công Brute-force và DoS/Spam API.

---

## ✨ Tính năng
* **Thuật toán Token Bucket:** Xử lý các đợt lưu lượng đột biến (bursty traffic) mượt mà mà vẫn đảm bảo tốc độ trung bình an toàn.
* **Theo dõi Client IP:** Lưu trữ bucket độc lập theo IP khách hàng hoặc API Key.
* **Chuẩn RFC 6585:** Tự động trả về mã lỗi HTTP `429 Too Many Requests` và header `Retry-After`.
* **Zero Dependencies:** Viết thuần Go 100%, không cần Redis hay thư viện ngoài.

---

## 🚀 Khởi chạy
```bash
go test -v ./...
go run ./cmd/example
```
