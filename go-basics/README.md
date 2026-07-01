# go-basics — Effective Go (go.dev)

Tài liệu học Go theo **đúng mục lục** [Effective Go](https://go.dev/doc/effective_go).

> Effective Go (2009) vẫn là nền idiomatic Go; bổ sung thêm [Tour](https://go.dev/tour/), [How to Write Go Code](https://go.dev/doc/code), generics/modules ở tài liệu mới hơn.

## Đọc trước

- [Tour of Go](https://go.dev/tour/)
- [How to Write Go Code](https://go.dev/doc/code)
- [Language Spec](https://go.dev/ref/spec)

---

## Mục lục (giống Effective Go)

| # | Mục | Tài liệu | Code ví dụ |
|---|-----|----------|------------|
| 1 | [Introduction](https://go.dev/doc/effective_go#introduction) | [doc/01-introduction.md](doc/01-introduction.md) | `01-introduction.go` |
| 2 | [Formatting](https://go.dev/doc/effective_go#formatting) | [doc/02-formatting.md](doc/02-formatting.md) | `02-formatting.go` |
| 3 | [Commentary](https://go.dev/doc/effective_go#commentary) | [doc/03-commentary.md](doc/03-commentary.md) | `03-commentary-names.go` |
| 4 | [Names](https://go.dev/doc/effective_go#names) | [doc/04-names.md](doc/04-names.md) | `03-commentary-names.go` |
| 5 | [Semicolons](https://go.dev/doc/effective_go#semicolons) | [doc/05-semicolons.md](doc/05-semicolons.md) | `04-semicolons.go` |
| 6 | [Control structures](https://go.dev/doc/effective_go#control_structures) | [doc/06-control-structures.md](doc/06-control-structures.md) | `05-control-if.go` |
| 7 | [Functions](https://go.dev/doc/effective_go#functions) | [doc/07-functions.md](doc/07-functions.md) | `08-functions.go` |
| 8 | [Data](https://go.dev/doc/effective_go#data) | [doc/08-data.md](doc/08-data.md) | `10-data-new-make.go` |
| 9 | [Initialization](https://go.dev/doc/effective_go#initialization) | [doc/09-initialization.md](doc/09-initialization.md) | `15-init-constants-variables.go` |
| 10 | [Methods](https://go.dev/doc/effective_go#methods) | [doc/10-methods.md](doc/10-methods.md) | `16-methods.go` |
| 11 | [Interfaces and other types](https://go.dev/doc/effective_go#interfaces) | [doc/11-interfaces-and-other-types.md](doc/11-interfaces-and-other-types.md) | `15-init-constants-variables.go` |
| 12 | [The blank identifier](https://go.dev/doc/effective_go#blank) | [doc/12-blank-identifier.md](doc/12-blank-identifier.md) | `19-blank-identifier.go` |
| 13 | [Embedding](https://go.dev/doc/effective_go#embedding) | [doc/13-embedding.md](doc/13-embedding.md) | `18-structs-embedding.go` |
| 14 | [Concurrency](https://go.dev/doc/effective_go#concurrency) | [doc/14-concurrency.md](doc/14-concurrency.md) | — |
| 15 | [Errors](https://go.dev/doc/effective_go#errors) | [doc/15-errors.md](doc/15-errors.md) | — |
| 16 | [A web server](https://go.dev/doc/effective_go#a_web_server) | [doc/16-a-web-server.md](doc/16-a-web-server.md) | — |

---

## Cấu trúc thư mục

```
go-basics/
├── README.md          ← mục lục (file này)
├── doc/               ← tài liệu từng đầu mục (giống Effective Go)
│   ├── 01-introduction.md
│   ├── 06-control-structures.md
│   └── ...
└── *.go               ← code ví dụ (package gobasics)
```

## Chạy test

```bash
go test ./go-basics/... -v
```

## Phần nâng cao

Interfaces · Concurrency · Errors → [`advance-go/`](../advance-go/)
