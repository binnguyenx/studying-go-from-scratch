# Go Learning Roadmap — 20 Days to Intern

> Target: Backend Engineering with Go (API, Concurrency, DB, Production)

**Daily study:** each day has a Go file with theory + exercise + hint in comments — open `day01/day01.go` … `day20/day20.go`. Run `go run ./day01` from the repo root.

---

## Day-by-Day Schedule

### Week 1 — Core Language Fundamentals

| Day | Topic | Content |
|-----|-------|---------|
| Day 1 | Setup + Basics | Install Go, VSCode extension, variables, types (`string`, `int`, `float64`, `bool`), `fmt.Println` |
| Day 2 | Control Flow + Functions | `if`, `switch`, `for` (Go has no `while`), function basics, named returns |
| Day 3 | Multiple Returns + Pointers | Multiple return values, `*` and `&`, pass by value vs pass by reference |
| Day 4 | Arrays + Slices | `[]int{}`, `len`, `cap`, `append`, `copy`, slice tricks |
| Day 5 | Maps + Structs | `map[string]int{}`, struct definition, struct methods, embedded structs |
| Day 6 | Interfaces | Implicit interface, duck typing, `interface{}` (any), type assertion |
| Day 7 | Review + Mini Exercise | Rewrite all concepts from memory, solve small problems |

---

### Week 2 — Error Handling + Packages + Concurrency

| Day | Topic | Content |
|-----|-------|---------|
| Day 8 | Error Handling | `if err != nil`, custom errors, `fmt.Errorf`, `errors.Is`, `errors.As`, wrapping errors |
| Day 9 | Packages + Modules | `go mod init`, `go.mod`, `import`, project structure, visibility (exported vs unexported) |
| Day 10 | Goroutines | `go func()`, goroutine lifecycle, `sync.WaitGroup` |
| Day 11 | Channels | `make(chan int)`, buffered vs unbuffered, `range` over channel, closing channel |
| Day 12 | Select + Mutex | `select` statement, `sync.Mutex`, race condition, deadlock |
| Day 13 | Context | `context.Background()`, `context.WithTimeout`, `context.WithCancel`, request lifecycle — [How to use Contexts in Go](https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go) |
| Day 14 | Concurrency Patterns | Worker pool, producer-consumer, fan-out fan-in |

---

### Week 3 — Backend Engineering

| Day | Topic | Content |
|-----|-------|---------|
| Day 15 | HTTP Server | `net/http`, `http.HandleFunc`, `http.ListenAndServe`, JSON encode/decode |
| Day 16 | REST API + Middleware | Router (`chi` or `gin`), middleware, path params, query params — [Middleware & Roundtrippers in Go](https://dev.to/calvinmclean/middleware-and-roundtrippers-in-go-30pa) |
| Day 17 | Testing | `go test`, table-driven tests, `testify`, `httptest` — see [Testing resources](#testing-1) |
| Day 18 | PostgreSQL + GORM | Connect DB, GORM models, CRUD, migrations, transactions |
| Day 19 | Auth + JWT | `golang-jwt/jwt`, bcrypt password hashing, middleware auth |
| Day 20 | Production Patterns | Structured logging, env config, graceful shutdown, Dockerfile, profiling with `pprof` |

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
└── for (the only loop in Go — replaces while too)

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
└── [5]int{} — fixed size, rarely used directly

Slices
├── []int{} — dynamic
├── make([]int, len, cap)
├── append, copy
├── Slice of slice
└── Pitfall: shared underlying array

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
├── Implicit implementation (no "implements" keyword)
├── Duck typing
├── Interface composition
├── Empty interface: interface{} / any
├── Type assertion: val.(Type)
└── Type switch

Key interfaces to know
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
Standard project layout
├── cmd/           — entry points (main.go)
├── internal/      — private packages
├── pkg/           — public packages
├── api/           — API definitions
└── go.mod / go.sum

Modules
├── go mod init <module-name>
├── go get <package>
├── go mod tidy
└── go build / go run
```

### 6. Concurrency (most important)

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

Context — [Tutorial](https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go)
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

Frameworks (pick one)
├── gin-gonic/gin — most popular
├── go-chi/chi — lightweight
└── labstack/echo

Middleware — [Middleware & Roundtrippers in Go](https://dev.to/calvinmclean/middleware-and-roundtrippers-in-go-30pa)
├── Logging
├── Auth (JWT)
├── CORS
└── Rate limiting
```

### 8. Testing

```
Standard library — [pkg.go.dev/testing](https://pkg.go.dev/testing)
├── testing.T
├── t.Run() — subtests
├── t.Parallel()
├── Table-driven tests
└── go test ./...

Useful packages
├── testify/assert
├── testify/mock
└── httptest.NewRecorder() — [HTTP test guide](https://speedscale.com/blog/testing-golang-with-httptest/)

Types
├── Unit test
├── Integration test
└── Benchmark: testing.B — [Benchmarking guide](https://betterstack.com/community/guides/scaling-go/golang-benchmarking/)

Test Coverage — [go.dev/doc/build-cover](https://go.dev/doc/build-cover)
```

### 9. Database

```
SQL
├── database/sql (stdlib)
├── lib/pq or pgx (PostgreSQL driver)
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

Profiling with pprof
├── net/http/pprof
└── go tool pprof

Dockerfile
├── Multi-stage build
└── Distroless / Alpine image — [Docker Compose docs](https://docs.docker.com/compose/)

Tracing / Monitoring (nice to have)
├── OpenTelemetry
└── Prometheus metrics
```

---

## Resources

### Official Go

| Resource | Link |
|----------|------|
| Official docs | [go.dev/doc](https://go.dev/doc) |
| Tour of Go | [go.dev/tour](https://go.dev/tour) |
| Effective Go | [go.dev/doc/effective_go](https://go.dev/doc/effective_go) |
| Go Playground | [go.dev/play](https://go.dev/play) |
| Go by Example | [gobyexample.com](https://gobyexample.com) |

### Go Best Practices (from company)

| Resource | Link |
|----------|------|
| Go Proverbs | [go-proverbs.github.io](https://go-proverbs.github.io/) |
| 50 Shades of Go — Traps & Gotchas | [golang50shades.com](https://golang50shades.com/) |

### Testing

| Resource | Link |
|----------|------|
| testing package | [pkg.go.dev/testing](https://pkg.go.dev/testing) |
| Benchmarking in Go | [betterstack.com](https://betterstack.com/community/guides/scaling-go/golang-benchmarking/) |
| HTTP test package (optional) | [speedscale.com](https://speedscale.com/blog/testing-golang-with-httptest/) |
| Middleware & Roundtrippers (optional) | [dev.to](https://dev.to/calvinmclean/middleware-and-roundtrippers-in-go-30pa) |
| Test Coverage (optional) | [go.dev/doc/build-cover](https://go.dev/doc/build-cover) |

### Concurrency & Internals

| Resource | Link |
|----------|------|
| How to use Contexts in Go | [digitalocean.com](https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go) |
| Profiling with pprof (video 1) | [youtu.be/N3PWzBeLX2M](https://youtu.be/N3PWzBeLX2M?si=brzhdBvpu82ill77) |
| Profiling with pprof (video 2) | [youtu.be/nok0aYiGiYA](https://youtu.be/nok0aYiGiYA?si=f7SNdeGv8iEEw3km) |

### DevOps & Tooling

| Resource | Link |
|----------|------|
| Docker Compose | [docs.docker.com/compose](https://docs.docker.com/compose/) |
| Git guide | [rogerdudler.github.io/git-guide](https://rogerdudler.github.io/git-guide/) |
| Git branching model | [nvie.com](https://nvie.com/posts/a-successful-git-branching-model/) |

> **Also recommended:** Get familiar with basic Linux commands — `ssh`, `vim`, `nano`, `sudo`, `grep`, `tail`, `find`

---

## Projects (in order)

1. **Todo API** — CRUD, JSON, Routing, Middleware
2. **Auth Service** — PostgreSQL, JWT, bcrypt
3. **URL Shortener** — Redis, short-lived keys
4. **Distributed Job Queue** — Goroutines, Channels, Worker pool
