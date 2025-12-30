
# Go Distributed Rate Limiter

A **high-performance, distributed rate limiting service** built in **Go**, designed to protect APIs from excessive traffic using the **Token Bucket algorithm** with **Redis-backed shared state**.

This project demonstrates **real-world backend engineering concepts** such as distributed coordination, concurrency safety, middleware design, and scalable traffic control.

---

## ✨ Features

* Token Bucket rate limiting algorithm
* Redis-backed distributed state (multi-instance safe)
* Per-client rate limiting (IP / API key)
* HTTP middleware for easy integration
* Configurable capacity and refill rate
* Lightweight and production-oriented design

---

## 🧱 Architecture Overview

```
Client
  ↓
Rate Limiter Service (Go)
  ↓
Redis (shared state)
```

Redis ensures consistent rate limits even when the service is horizontally scaled.

---

## 📁 Project Structure

```
go-rate-limiter/
├── cmd/
│   └── server/
│       └── main.go        # Application entry point
├── internal/
│   ├── api/               # HTTP handlers
│   ├── limiter/           # Rate limiting algorithms
│   │   └── token_bucket.go
│   ├── middleware/        # HTTP middleware
│   ├── storage/           # Redis abstraction layer
│   └── config/            # Configuration management
├── go.mod
└── README.md
```

This structure follows **Go best practices**:

* `cmd/` for executables
* `internal/` to enforce encapsulation
* Clear separation of concerns

---

## ⚙️ How It Works

1. Each client request is identified using an IP address or API key
2. The Token Bucket algorithm determines if the request is allowed
3. Token state is stored and synchronized using Redis
4. Requests exceeding the limit receive an HTTP `429 Too Many Requests` response
5. Middleware allows seamless API protection

---

## 🛠️ Installation

### Prerequisites

* Go **1.21+**
* Redis

### Clone and Setup

```bash
git clone https://github.com/JagTheFriend/Go-Rate-Limiter.git
cd Go-Rate-Limiter
go mod tidy
```

---

## ▶️ Running the Service

Start Redis:

```bash
docker run -p 6379:6379 redis
```

Run the server:

```bash
go run cmd/server/main.go
```

The service will start on:

```
http://localhost:8080
```

---

## 🔌 API Usage

### Check Rate Limit

```
GET /allow?key=user123
```

### Sample Response

```json
{
  "allowed": true,
  "remaining": 7
}
```

### Rate Limited Response

```
HTTP/1.1 429 Too Many Requests
Rate limit exceeded
```

---

## 🧠 Key Concepts Demonstrated

* Token Bucket algorithm
* Distributed systems using Redis
* Concurrency-safe request handling
* Middleware-based API protection
* Clean and maintainable Go architecture

---

## 📈 Possible Enhancements

* Sliding Window rate limiting
* Redis Lua scripts for atomic updates
* Per-route and per-method limits
* Prometheus-compatible metrics
* Docker Compose setup
* Configurable rate limits via YAML/ENV
