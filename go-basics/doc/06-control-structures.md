# Control structures

> **Nguồn:** [Control structures](https://go.dev/doc/effective_go#control_structures) · [Effective Go](https://go.dev/doc/effective_go)

Khác C: **không** `do`/`while`; `for` hơi tổng quát; `switch` linh hoạt hơn; `if`/`switch`/`for` có thể có **init statement**; `break`/`continue` nhận **label**; có **type switch** và `select` (xem Concurrency).

**Syntax:** không ngoặc quanh điều kiện; thân luôn trong `{ }`.

---

## If

```go
if x > 0 {
    return y
}
```

Init + check (rất phổ biến với error):

```go
if err := file.Chmod(0664); err != nil {
    return err
}
```

**Style:** nếu nhánh `if` kết thúc bằng `return`/`break`/`continue`/`goto` → **bỏ `else` thừa**; xử lý lỗi xong, luồng chính chạy xuống dưới.

**Code:** [`../05-control-if.go`](../05-control-if.go) — `Sign`, `ErrRedeclarationDemo`

↳ [go.dev — If](https://go.dev/doc/effective_go#if)

---

## Redeclaration and reassignment

```go
f, err := os.Open(name)
// ...
d, err := f.Stat()  // err: gán lại, không khai báo mới
```

`:=` hợp lệ khi **cùng scope**, **ít nhất một biến mới**, và value assignable.

**Code:** [`../05-control-if.go`](../05-control-if.go)

↳ [go.dev — Redeclaration](https://go.dev/doc/effective_go#redeclaration)

---

## For

Ba dạng:

```go
for init; condition; post { }  // như C
for condition { }              // như while
for { }                        // vô hạn — break/return
```

`range` trên slice, array, string, map, channel:

```go
for key, value := range m { }
for _, value := range array { }  // bỏ index bằng _
```

String `range`: index = **byte**, value = **rune** (UTF-8). Byte lỗi → `U+FFFD`.

Reverse song song (không có comma operator):

```go
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

**Code:** [`../06-control-for.go`](../06-control-for.go)

↳ [go.dev — For](https://go.dev/doc/effective_go#for)

---

## Switch

Tagless switch (thay if-else chain):

```go
switch {
case x < 0:
case x == 0:
default:
}
```

Không fallthrough (trừ `fallthrough`); case nhiều giá trị: `case 1, 2, 3:`

**Code:** [`../07-control-switch.go`](../07-control-switch.go) — `FizzBuzz`, `Unhex`

↳ [go.dev — Switch](https://go.dev/doc/effective_go#switch)

---

## Type switch

```go
switch t := t.(type) {
case bool:
case int:
default:
}
```

**Code:** [`../07-control-switch.go`](../07-control-switch.go) — `TypeSwitchDescribe`

↳ [go.dev — Type switch](https://go.dev/doc/effective_go#type_switch)
