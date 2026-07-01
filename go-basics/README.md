# go-basics — study path (Effective Go)

Follow **[Effective Go](https://go.dev/doc/effective_go)** on go.dev in this order. Each file maps to a doc section; read the link first, then the code.

Also read first (go.dev): [Tour of Go](https://go.dev/tour/) · [How to Write Go Code](https://go.dev/doc/code) · [Language Spec](https://go.dev/ref/spec)

| # | File | Effective Go section |
|---|------|----------------------|
| 01 | [01-introduction.go](01-introduction.go) | [Introduction](https://go.dev/doc/effective_go#introduction) |
| 02 | [02-formatting.go](02-formatting.go) | [Formatting](https://go.dev/doc/effective_go#formatting) |
| 03 | [03-commentary-names.go](03-commentary-names.go) | [Commentary](https://go.dev/doc/effective_go#commentary) · [Names](https://go.dev/doc/effective_go#names) |
| 04 | [04-semicolons.go](04-semicolons.go) | [Semicolons](https://go.dev/doc/effective_go#semicolons) |
| 05 | [05-control-if.go](05-control-if.go) | [Control structures → If](https://go.dev/doc/effective_go#if) · [Redeclaration](https://go.dev/doc/effective_go#redeclaration) |
| 06 | [06-control-for.go](06-control-for.go) | [For](https://go.dev/doc/effective_go#for) |
| 07 | [07-control-switch.go](07-control-switch.go) | [Switch](https://go.dev/doc/effective_go#switch) · [Type switch](https://go.dev/doc/effective_go#type_switch) |
| 08 | [08-functions.go](08-functions.go) | [Functions](https://go.dev/doc/effective_go#functions) · [Multiple returns](https://go.dev/doc/effective_go#multiple-returns) |
| 09 | [09-defer.go](09-defer.go) | [Defer](https://go.dev/doc/effective_go#defer) |
| 10 | [10-data-new-make.go](10-data-new-make.go) | [Data → new](https://go.dev/doc/effective_go#allocation_new) · [make](https://go.dev/doc/effective_go#allocation_make) |
| 11 | [11-arrays.go](11-arrays.go) | [Arrays](https://go.dev/doc/effective_go#arrays) |
| 12 | [12-slices.go](12-slices.go) | [Slices](https://go.dev/doc/effective_go#slices) · [Append](https://go.dev/doc/effective_go#append) |
| 13 | [13-maps.go](13-maps.go) | [Maps](https://go.dev/doc/effective_go#maps) |
| 14 | [14-printing.go](14-printing.go) | [Printing](https://go.dev/doc/effective_go#printing) |
| 15 | [15-init-constants-variables.go](15-init-constants-variables.go) | [Initialization](https://go.dev/doc/effective_go#initialization) |
| 16 | [16-methods.go](16-methods.go) | [Methods](https://go.dev/doc/effective_go#methods) · [Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values) |
| 17 | [17-pointers.go](17-pointers.go) | (with Methods — receivers & `*T`) |
| 18 | [18-structs-embedding.go](18-structs-embedding.go) | [Constructors & composite literals](https://go.dev/doc/effective_go#composite_literals) · [Embedding](https://go.dev/doc/effective_go#embedding) |
| 19 | [19-blank-identifier.go](19-blank-identifier.go) | [The blank identifier](https://go.dev/doc/effective_go#blank) |
| 20 | [20-packages-imports.go](20-packages-imports.go) | Packages (see [How to Write Go Code](https://go.dev/doc/code)) |

**Continue on go.dev (see `advance-go/` in repo):** [Interfaces](https://go.dev/doc/effective_go#interfaces) · [Concurrency](https://go.dev/doc/effective_go#concurrency) · [Errors](https://go.dev/doc/effective_go#errors)

```bash
go test ./go-basics/... -v
```
