# Initialization

> **Nguồn:** [Initialization](https://go.dev/doc/effective_go#initialization) · [Effective Go](https://go.dev/doc/effective_go)

Đọc bản gốc trên go.dev trước, rồi xem code ví dụ trong repo (nếu có).

## Constants

Compile-time; iota cho enum.

**Code:** [`../15-init-constants-variables.go`](../15-init-constants-variables.go)

↳ [go.dev — Constants](https://go.dev/doc/effective_go#constants)

## Variables

Init lúc runtime; `var x = expr`.

**Code:** [`../15-init-constants-variables.go`](../15-init-constants-variables.go)

↳ [go.dev — Variables](https://go.dev/doc/effective_go#variables)

## The init function

`func init()` mỗi file; chạy sau import.

↳ [go.dev — The init function](https://go.dev/doc/effective_go#init)
