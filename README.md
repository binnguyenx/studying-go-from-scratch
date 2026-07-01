# studying-go-from-scratch

Go learning repo aligned with **Self Made Engineer** and company onboarding resources.

> **Workflow:** theory first (course video + links below) → open the matching file in `advance-go/` → code exercises in that file or add `*_test.go` → `go test ./advance-go/...`

---

## Repo layout

| Folder | Course | Status |
|--------|--------|--------|
| [`advance-go/`](advance-go/) | **Advanced Go** — entry track (essentials before other courses) | Active — one `.go` file per module |
| [`go-basics/`](go-basics/) | **Basic Go syntax** — variables, control flow, collections, pointers | Active — one file per topic |
| `backend-engineering-with-go/` | Backend Engineering with Go — APIs, deploy | Planned |
| `microservices-with-go/` | Microservices — Kubernetes, distributed systems | Planned |
| `order-management-system/` | Building an Order Management System — CockroachDB, Stripe, RabbitMQ, Docker | Planned |

---

## `go-basics/` — Basic Go syntax

Study **before** or alongside `advance-go/`. Each file is `package gobasics` with syntax notes in the header + small demo functions.

| File | Topic |
|------|--------|
| [`01-variables-constants.go`](go-basics/01-variables-constants.go) | `var`, `:=`, `const`, `iota` |
| [`02-types-zero-values.go`](go-basics/02-types-zero-values.go) | Built-in types, zero values |
| [`03-operators.go`](go-basics/03-operators.go) | Arithmetic, comparison, logical, bitwise |
| [`04-if-switch.go`](go-basics/04-if-switch.go) | `if`, `switch`, tagless switch |
| [`05-for-loops.go`](go-basics/05-for-loops.go) | `for` — 3 forms, `range`, string runes ([Effective Go §For](https://go.dev/doc/effective_go#for)) |
| [`06-functions.go`](go-basics/06-functions.go) | Functions, multiple returns, variadic |
| [`07-arrays.go`](go-basics/07-arrays.go) | Fixed-size arrays |
| [`08-slices.go`](go-basics/08-slices.go) | Slices, `append`, `copy` |
| [`09-maps.go`](go-basics/09-maps.go) | Maps, `ok` idiom |
| [`10-structs.go`](go-basics/10-structs.go) | Structs, embedding |
| [`11-methods.go`](go-basics/11-methods.go) | Value vs pointer receivers |
| [`12-pointers.go`](go-basics/12-pointers.go) | `&`, `*`, `nil` |
| [`13-strings-runes.go`](go-basics/13-strings-runes.go) | Strings, UTF-8, `range` runes |
| [`14-defer.go`](go-basics/14-defer.go) | `defer` LIFO |
| [`15-type-conversion.go`](go-basics/15-type-conversion.go) | Explicit conversions |
| [`16-range.go`](go-basics/16-range.go) | `range` on slice, map, string |
| [`17-fmt-printing.go`](go-basics/17-fmt-printing.go) | `Printf` verbs (see also `fmt` pkg) |
| [`18-packages-imports.go`](go-basics/18-packages-imports.go) | Packages, exports, imports |

```bash
go test ./go-basics/... -v
```

---

## `advance-go/` — Advanced Go modules

Follow the course sidebar in order. Each file is `package advancego` with **theory + exercise + hint** in the header comments.

| # | File | Topic |
|---|------|--------|
| 1 | [`introduction.go`](advance-go/introduction.go) | Introduction |
| 2 | [`codealong-platform.go`](advance-go/codealong-platform.go) | CodeAlong the new platform experience |
| 3 | [`effective-error-handling.go`](advance-go/effective-error-handling.go) | Effective Error Handling |
| 4 | [`interfaces.go`](advance-go/interfaces.go) | Interfaces |
| 5 | [`testing.go`](advance-go/testing.go) | Testing |
| 6 | [`pointers.go`](advance-go/pointers.go) | Pointers |
| 7 | [`goroutines.go`](advance-go/goroutines.go) | Goroutines |
| 8 | [`context-and-timeouts.go`](advance-go/context-and-timeouts.go) | Context and Timeouts |
| 9 | [`concurrency-with-channels.go`](advance-go/concurrency-with-channels.go) | Concurrency with Channels |
| 10 | [`maps.go`](advance-go/maps.go) | Maps |
| 11 | [`capstone-project.go`](advance-go/capstone-project.go) | Capstone Project (Exercise) |
| 12 | [`project-solution.go`](advance-go/project-solution.go) | Project solution |
| 13 | [`map-concurrency-mutexes.go`](advance-go/map-concurrency-mutexes.go) | Map Concurrency & Mutexes |
| 14 | [`where-to-go-next.go`](advance-go/where-to-go-next.go) | Where to go next? |

```bash
# From repo root — compile the package
go build ./advance-go/...

# Run tests after you add them
go test ./advance-go/... -race
```

---

## Next courses (after Advanced Go)

| Track | Focus |
|-------|--------|
| **Backend Engineering with Go** | Full backend web dev — build and deploy APIs |
| **Microservices with Go** | Scalable services with Kubernetes |
| **Order Management System** | Microservices capstone — CockroachDB, Stripe, RabbitMQ, Docker |

Add a new top-level folder per track when you start that course (same pattern: one file per module).

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

### Go best practices (company)

| Resource | Link |
|----------|------|
| Go Proverbs | [go-proverbs.github.io](https://go-proverbs.github.io/) |
| 50 Shades of Go | [golang50shades.com](https://golang50shades.com/) |

### Testing

| Resource | Link |
|----------|------|
| testing package | [pkg.go.dev/testing](https://pkg.go.dev/testing) |
| Benchmarking | [betterstack.com](https://betterstack.com/community/guides/scaling-go/golang-benchmarking/) |
| httptest (optional) | [speedscale.com](https://speedscale.com/blog/testing-golang-with-httptest/) |
| Middleware & Roundtrippers (optional) | [dev.to](https://dev.to/calvinmclean/middleware-and-roundtrippers-in-go-30pa) |
| Test coverage (optional) | [go.dev/doc/build-cover](https://go.dev/doc/build-cover) |

### Concurrency & internals

| Resource | Link |
|----------|------|
| Contexts in Go | [digitalocean.com](https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go) |
| pprof (video 1) | [youtu.be/N3PWzBeLX2M](https://youtu.be/N3PWzBeLX2M?si=brzhdBvpu82ill77) |
| pprof (video 2) | [youtu.be/nok0aYiGiYA](https://youtu.be/nok0aYiGiYA?si=f7SNdeGv8iEEw3km) |

### DevOps & tooling

| Resource | Link |
|----------|------|
| Docker Compose | [docs.docker.com/compose](https://docs.docker.com/compose/) |
| Git guide | [rogerdudler.github.io/git-guide](https://rogerdudler.github.io/git-guide/) |
| Git branching model | [nvie.com](https://nvie.com/posts/a-successful-git-branching-model/) |

> **Also:** basic Linux — `ssh`, `vim`, `nano`, `sudo`, `grep`, `tail`, `find`
