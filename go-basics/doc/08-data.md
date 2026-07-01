# Data

> **Nguồn:** [Data](https://go.dev/doc/effective_go#data) · [Effective Go](https://go.dev/doc/effective_go)

Hai primitive cấp phát: **`new`** và **`make`** — khác nhau, dùng cho kiểu khác nhau.

---

## Allocation with `new`

- `new(T)` → `*T`, memory **zeroed**
- Không “constructor”; zero value nhiều khi đã dùng được (`bytes.Buffer`, `sync.Mutex`)

**Code:** [`../10-data-new-make.go`](../10-data-new-make.go)

↳ [go.dev — new](https://go.dev/doc/effective_go#allocation_new)

---

## Constructors and composite literals

```go
f := File{fd: fd, name: name}
return &File{fd: fd, name: name}
```

`&T{}` tạo instance mới mỗi lần. `new(T)` và `&T{}` tương đương với struct rỗng field.

**Code:** [`../18-structs-embedding.go`](../18-structs-embedding.go)

↳ [go.dev — composite literals](https://go.dev/doc/effective_go#composite_literals)

---

## Allocation with `make`

- Chỉ **slice, map, channel**
- Trả **value** (không phải pointer), **initialized**
- `make([]int, 10, 100)` — len 10, cap 100

`new([]int)` → `*[]int` trỏ tới **nil slice** (hiếm dùng).

**Code:** [`../10-data-new-make.go`](../10-data-new-make.go), [`../12-slices.go`](../12-slices.go)

↳ [go.dev — make](https://go.dev/doc/effective_go#allocation_make)

---

## Arrays

- **Value type** — assign = copy toàn bộ
- Pass vào hàm = copy (trừ khi truyền pointer)
- `[10]int` và `[20]int` là **hai kiểu khác nhau**
- Thực tế: dùng **slice** thay array

**Code:** [`../11-arrays.go`](../11-arrays.go)

↳ [go.dev — Arrays](https://go.dev/doc/effective_go#arrays)

---

## Slices

- View lên array; assign slice = **cùng backing** (thường)
- Hàm nhận `[]T` → đổi phần tử caller thấy
- `len`, `cap`, `append`, `copy`

**Code:** [`../12-slices.go`](../12-slices.go)

↳ [go.dev — Slices](https://go.dev/doc/effective_go#slices)

---

## Two-dimensional slices

```go
type LinesOfText [][]byte
```

Mỗi hàng có thể **độ dài khác nhau**. Cấp phát từng dòng hoặc một block rồi slice.

↳ [go.dev — Two-dimensional slices](https://go.dev/doc/effective_go#two-dimensional-slices)

---

## Maps

- Key phải **comparable**
- Reference type — pass map, đổi trong hàm caller thấy
- **Không** concurrent-safe

**Code:** [`../13-maps.go`](../13-maps.go)

↳ [go.dev — Maps](https://go.dev/doc/effective_go#maps)

---

## Printing

`fmt` — verbs, `Stringer`, spacing rules.

**Code:** [`../14-printing.go`](../14-printing.go)

↳ [go.dev — Printing](https://go.dev/doc/effective_go#printing)

---

## Append

Built-in `append`; **phải gán lại** biến slice. `append(s, y...)`.

**Code:** [`../12-slices.go`](../12-slices.go)

↳ [go.dev — Append](https://go.dev/doc/effective_go#append)
