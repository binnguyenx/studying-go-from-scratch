# Functions

> **Nguồn:** [Functions](https://go.dev/doc/effective_go#functions) · [Effective Go](https://go.dev/doc/effective_go)

Đọc bản gốc trên go.dev trước, rồi xem code ví dụ trong repo (nếu có).

## Multiple return values

Hàm trả `(n int, err error)` — idiom Go.

**Code:** [`../08-functions.go`](../08-functions.go)

↳ [go.dev — Multiple return values](https://go.dev/doc/effective_go#multiple-returns)

## Named result parameters

Named return + naked return khi logic ngắn.

**Code:** [`../08-functions.go`](../08-functions.go)

↳ [go.dev — Named result parameters](https://go.dev/doc/effective_go#named-result-parameters)

## Defer

Chạy trước khi hàm return; LIFO.

**Code:** [`../09-defer.go`](../09-defer.go)

↳ [go.dev — Defer](https://go.dev/doc/effective_go#defer)
