# Go Learning Roadmap — 20 Days to Intern

> Target: Backend Engineering với Go (API, Concurrency, DB, Production)

---

## Day-by-Day Schedule

### Week 1 — Core Language Fundamentals

| Day | Chủ đề | Nội dung |
|-----|--------|----------|
| Day 1 | Setup + Basics | Cài Go, VSCode extension, variables, types (`string`, `int`, `float64`, `bool`), `fmt.Println` |
| Day 2 | Control Flow + Functions | `if`, `switch`, `for` (Go không có `while`), function basics, named returns |
| Day 3 | Multiple Returns + Pointers | Multiple return values, `*` và `&`, pass by value vs pass by reference |
| Day 4 | Arrays + Slices | `[]int{}`, `len`, `cap`, `append`, `copy`, slice tricks |
| Day 5 | Maps + Structs | `map[string]int{}`, struct definition, struct methods, embedded structs |
| Day 6 | Interfaces | Implicit interface, duck typing, `interface{}` (any), type assertion |
| Day 7 | Review + Mini Exercise | Viết lại các concept, làm bài tập tổng hợp |

---

### Week 2 — Error Handling + Packages + Concurrency

| Day | Chủ đề | Nội dung |
|-----|--------|----------|
| Day 8 | Error Handling | `if err != nil`, custom errors, `fmt.Errorf`, `errors.Is`, `errors.As`, wrapping errors |
| Day 9 | Packages + Modules | `go mod init`, `go.mod`, `import`, project structure, visibility (exported vs unexported) |
| Day 10 | Goroutines | `go func()`, goroutine lifecycle, `sync.WaitGroup` |
| Day 11 | Channels | `make(chan int)`, buffered vs unbuffered, `range` over channel, closing channel |
| Day 12 | Select + Mutex | `select` statement, `sync.Mutex`, race condition, deadlock |
| Day 13 | Context | `context.Background()`, `context.WithTimeout`, `context.WithCancel`, request lifecycle |
| Day 14 | Concurrency Patterns | Worker pool, producer-consumer, fan-out fan-in |

---

### Week 3 — Backend Engineering

| Day | Chủ đề | Nội dung |
|-----|--------|----------|
| Day 15 | HTTP Server | `net/http`, `http.HandleFunc`, `http.ListenAndServe`, JSON encode/decode |
| Day 16 | REST API | Router (dùng `chi` hoặc `gin`), middleware, path params, query params |
| Day 17 | Testing | `go test`, `testing.T`, table-driven tests, `testify`, mock |
| Day 18 | PostgreSQL + GORM | Connect DB, GORM models, CRUD, migrations, transactions |
| Day 19 | Auth + JWT | `golang-jwt/jwt`, bcrypt password hashing, middleware auth |
| Day 20 | Production Patterns | Structured logging (`zerolog`/`zap`), env config, graceful shutdown, Dockerfile |

---

## Full Knowledge Map

### 1. Language Core

```
Variables & Types
├── var, :=, const
├── Zero values
├── Type conversion
└── Type inference

Control Flow
├── if / else if / else
├── switch (no fallthrough by default)
└── for (là while duy nhất của Go)

Functions
├── Basic function
├── Multiple return values
├── Named return values
├── Variadic functions (...T)
├── First-class functions (pass func as arg)
├── Closures
└── defer

Pointers
├── &var (address of)
├── *ptr (dereference)
├── Pointer receivers
└── nil pointer
```

### 2. Data Structures

```
Arrays
└── [5]int{} — fixed size, rarely used

Slices
├── []int{} — dynamic
├── make([]int, len, cap)
├── append, copy
├── Slice of slice
└── Pitfalls: shared underlying array

Maps
├── map[K]V{}
├── make(map[K]V)
├── delete(m, key)
├── ok idiom: val, ok := m[key]
└── Not safe for concurrent use

Structs
├── type Foo struct {}
├── Anonymous fields (embedding)
├── Struct tags (json:"name")
└── Methods: func (f Foo) Method() {}
```

### 3. Interfaces & OOP Mindset

```
Interfaces
├── Implicit implementation (no "implements")
├── Duck typing
├── Interface composition
├── Empty interface: interface{} / any
├── Type assertion: val.(Type)
└── Type switch

Common Interfaces biết là mạnh
├── error
├── io.Reader / io.Writer
├── fmt.Stringer
└── http.Handler
```

### 4. Error Handling

```
Error pattern
├── if err != nil { return err }
├── Sentinel errors: var ErrNotFound = errors.New(...)
├── Custom error type: type AppError struct {}
├── fmt.Errorf("context: %w", err)
├── errors.Is(err, target)
└── errors.As(err, &target)
```

### 5. Packages & Project Structure

```
Project layout (chuẩn)
├── cmd/           — entry points (main.go)
├── internal/      — private packages
├── pkg/           — public packages
├── api/           — API definitions
└── go.mod / go.sum

Modules
├── go mod init <module-name>
├── go get <package>
├── go tidy
└── go build / go run
```

### 6. Concurrency (quan trọng nhất)

```
Goroutines
├── go func() {}()
├── Lightweight thread (~2KB stack)
└── sync.WaitGroup

Channels
├── make(chan T)
├── make(chan T, bufferSize)
├── ch <- val (send)
├── val := <-ch (receive)
├── close(ch)
└── range ch

Sync Primitives
├── sync.Mutex / sync.RWMutex
├── sync.Once
├── sync.WaitGroup
└── sync/atomic

Select
├── select { case <-ch: }
├── default case (non-blocking)
└── timeout pattern

Patterns
├── Worker Pool
├── Fan-out / Fan-in
├── Producer-Consumer
├── Pipeline
└── Done channel

Context
├── context.Background()
├── context.TODO()
├── context.WithTimeout()
├── context.WithDeadline()
├── context.WithCancel()
└── ctx.Done(), ctx.Err()
```

### 7. HTTP + API

```
net/http
├── http.HandleFunc
├── http.Handler interface
├── http.Request (Body, Header, URL, Method)
├── http.ResponseWriter
└── http.ListenAndServe

JSON
├── json.Marshal / json.Unmarshal
├── json.Encoder / json.Decoder
└── struct tags: json:"field,omitempty"

Frameworks (chọn 1)
├── gin-gonic/gin — phổ biến nhất
├── go-chi/chi — lightweight
└── labstack/echo

Middleware
├── Logging
├── Auth (JWT)
├── CORS
└── Rate limiting
```

### 8. Testing

```
Standard library
├── testing.T
├── t.Run() — subtests
├── t.Parallel()
├── Table-driven tests
└── go test ./...

Packages hay dùng
├── testify/assert
├── testify/mock
└── httptest.NewRecorder()

Types
├── Unit test
├── Integration test
└── Benchmark: testing.B
```

### 9. Database

```
SQL
├── database/sql (stdlib)
├── lib/pq hoặc pgx (PostgreSQL driver)
└── sqlx

ORM
└── GORM
    ├── AutoMigrate
    ├── CRUD
    ├── Associations (HasMany, BelongsTo)
    ├── Transactions
    └── Hooks

Patterns
├── Repository pattern
└── Connection pooling
```

### 10. Production Engineering

```
Logging
├── log/slog (stdlib Go 1.21+)
├── zerolog
└── uber-go/zap

Config
├── os.Getenv
├── joho/godotenv (.env file)
└── viper

Graceful Shutdown
├── os.Signal
├── signal.Notify
└── http.Server.Shutdown(ctx)

Dockerfile cơ bản
├── Multi-stage build
└── Distroless / Alpine image

Tracing / Monitoring (nice to have)
├── OpenTelemetry
└── Prometheus metrics
```

---

## Resources

| Loại | Link |
|------|------|
| Official docs | https://go.dev/doc |
| Go by Example | https://gobyexample.com |
| Tour of Go | https://go.dev/tour |
| Effective Go | https://go.dev/doc/effective_go |
| Go Playground | https://go.dev/play |

---

## Projects (theo thứ tự)

1. **Todo API** — CRUD, JSON, Routing, Middleware
2. **Auth Service** — PostgreSQL, JWT, bcrypt
3. **URL Shortener** — Redis, short-lived keys
4. **Distributed Job Queue** — Goroutines, Channels, Worker pool
