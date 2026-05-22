# Daily exercises (end of each day)

**Per-day study file:** open `dayNN/dayNN.go` (theory + exercise + hint in comments). Implement in that folder (`main.go` extra files OK).

Complete after theory. Use hints before asking for full answers.

| Day | Topic | Exercise | Hint |
|-----|-------|----------|------|
| 1 | Setup + Basics | In `day01/day01.go`: warm-up (one `Printf` for `name`, `age`, `gpa`) + **Course roster** report (`Println` / `Printf` / `Sprintf`, all types `string`/`int`/`float64`/`bool`, padded table, no loops required). | See comment block “Main exercise” in `day01.go`; padding: `%-12s`, `%6.1f`, `%t`. |
| 2 | Control flow + Functions | Implement `FizzBuzz(n int) string` for one number, then print 1–30 from `main`. | `switch` or `if`; no `while`, only `for`. |
| 3 | Multiple returns + Pointers | `Divide(a, b int) (int, error)` returning quotient; `main` prints result or error. Add `Swap(a, b *int)` swapping values in `main`. | `errors.New("...")` for `b==0`; swap uses temp or arithmetic (pick one). |
| 4 | Arrays + Slices | `ReverseInPlace([]int) []int` reverses the **same** underlying data (no new slice). Unit-test with two cases. | Two indices `i`, `j` moving toward center. |
| 5 | Maps + Structs | `CountWords(s string) map[string]int` (words = split on spaces). Define `type Person struct { Name string; Age int }` and a method `Introduce() string`. | `strings.Fields`; `fmt.Sprintf` in method. |
| 6 | Interfaces | Define `type Speaker interface { Say() string }`. Implement it for two structs. Write `PrintSay(s Speaker)` and call with both. | No `implements` keyword — just add `Say() string` on each type. |
| 7 | Review | Mini CLI: read a line of space-separated ints, print sum and average (float). Handle bad input with message. | `bufio.Scanner`, `strconv.Atoi`, loop `for`. |
| 8 | Error handling | Create sentinel `ErrEmpty`. Function `ParsePositive(s string) (int, error)` returns wrapped error with `%w` when empty or non-positive. In `main`, demonstrate `errors.Is`. | `fmt.Errorf("parse %q: %w", s, ErrEmpty)`. |
| 9 | Packages + modules | `go mod init` module `learner/exercises`. Package `internal/mathx` with `Sum(ints []int) int`; `cmd/day09/main` imports and prints `Sum`. | Exported `Sum` capital S; `internal/` only importable from module root. |
| 10 | Goroutines | Fire 5 goroutines that each `time.Sleep` random ms, `WaitGroup` until all done, print “all done” once. | `wg.Add(1)` before `go`, `defer wg.Done()` inside goroutine. |
| 11 | Channels | `SquarePipe(nums []int) []int`: goroutine sends squares on channel; main collects slice. Close channel when producer finishes. | `go func(){ defer close(ch); for _, n := range nums { ... } }()`. |
| 12 | Select + Mutex | Bank account `Deposit` / `Withdraw` with `sync.Mutex`; run concurrent ops from multiple goroutines, assert final balance in test. | Lock around read-modify-write; `go test -race`. |
| 13 | Context | HTTP client GET with `context.WithTimeout` (2s) to slow URL or `http.NewRequestWithContext` to `https://httpbin.org/delay/5` — expect timeout error. | `ctx, cancel := context.WithTimeout(...); defer cancel()`. |
| 14 | Concurrency patterns | Worker pool: 3 workers consume job IDs 1–20 from channel; print each job once. | Separate `jobs` chan, `sync.WaitGroup` for workers, close `jobs` after enqueue. |
| 15 | HTTP + JSON | In-memory `map[string]User` + `net/http`: `POST /users` JSON body, `GET /users/{id}`. No framework required. | `json.NewDecoder(r.Body)`, `mux` with path trim or `strings.TrimPrefix`. |
| 16 | REST + middleware | Same as Day 15 or add `chi`/`gin`: logging middleware that logs method + path + duration. | Wrap `http.Handler` with `time.Since`; call `next.ServeHTTP`. |
| 17 | Testing | Table-driven tests for `IsPalindrome(s string) bool` (ASCII only). Add one benchmark. | Slice of struct `{name, in, want}`; `b.ResetTimer()` in benchmark loop. |
| 18 | DB | SQLite in-memory GORM: `User` model, migrate, create + find + update. | `gorm.io/driver/sqlite`, `dsn := "file::memory:?cache=shared"`. |
| 19 | Auth | Register endpoint hashing password with `bcrypt`; login returns signed JWT (HS256). Protect `GET /me` with middleware parsing `Authorization: Bearer`. | `jwt.NewWithClaims`, `ParseWithClaims`, check `errors.Is` for expiry. |
| 20 | Production | `http.Server` + graceful shutdown on `SIGINT`; expose `pprof` on debug mux or route. Read `PORT` from env default `8080`. | `signal.Notify`, `srv.Shutdown(ctx)`; `_ "net/http/pprof"`. |

## End of roadmap

Pick one project from `README.md` (Todo API → Auth → URL shortener → job queue) and track issues/tasks in your own notes.
