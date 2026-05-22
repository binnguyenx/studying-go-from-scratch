# Day 2 — Cập nhật lý thuyết (Control flow + Functions)

> Tóm tắt lý thuyết; bài tập và code mẫu: `day02.go`.

## Mục tiêu

- Điều kiện: `if` / `else`, `switch`.
- Lặp: chỉ có `for` (không có `while` riêng).
- Hàm: tham số, giá trị trả về, **named return** (tùy dùng).

## `if`

- Cú pháp cơ bản: `if cond { }` — ngoặc `()` **không** bắt buộc quanh điều kiện (khác C).
- Có thể gắn **câu lệnh khởi tạo** trước điều kiện: `if x := f(); x > 0 { ... }` — biến `x` chỉ tồn tại trong nhánh `if/else if/else` của khối đó.
- `else if`, `else` như thường.

## `switch`

- `switch x { case v1: ... case v2: ... default: ... }` — so sánh `x` với từng `case` (bằng, không cần `break` ở cuối case; **không rơi** xuống case kế — trừ khi ghi rõ `fallthrough`).
- **Tagless switch** (không có biểu thức sau `switch`): viết như chuỗi `if-else` sạch hơn:

```go
switch {
case t < 0:
    // ...
case t == 0:
    // ...
default:
    // ...
}
```

- Nhiều giá trị một case: `case 1, 2, 3:`.

## `for` — mọi vòng lặp đều là `for`

| Dạng | Ý nghĩa |
|------|--------|
| `for i := 0; i < n; i++ { }` | Giống C: init; điều kiện; post |
| `for cond { }` | Giống **while**: chỉ còn điều kiện |
| `for { }` | Vòng lặp vô hạn — cần `break` / `return` |
| `for i, v := range xs { }` | Duyệt slice, string, map, channel (ngày sau dùng nhiều hơn) |

- `break` thoát vòng gần nhất; `continue` sang lần lặp tiếp theo.
- Nhãn + `break label` khi cần thoát khỏi vòng ngoài (ít dùng ngày đầu).

## Hàm (functions)

- Khai báo: `func name(param T, other U) returnType { }`.
- Nhiều tham số cùng kiểu: `func add(a, b int) int`.
- **Nhiều giá trị trả về** (ngày 3 sẽ đi sâu với `error`): `func f() (int, bool)`.
- **Named return**: đặt tên cho biến kết quả trong chữ ký; `return` trần sẽ trả các giá trị hiện tại của chúng (dễ đọc nhưng dễ nhầm nếu logic phức tạp — dùng khi thật sự giúp rõ ý).

```go
func bounds(a, b int) (min, max int) {
    if a < b {
        min, max = a, b
    } else {
        min, max = b, a
    }
    return // "naked return" — trả min, max đã gán
}
```

- Tham số biến đổi (variadic): `func sum(xs ...int)` — gọi `sum(1,2,3)` hoặc truyền slice đã có: `sum(slice...)`.

## Go không có `while`

- Dùng `for condition { }` thay cho while.

## Đọc thêm

- [Tour of Go — For, If, Switch](https://go.dev/tour/list)
- [Effective Go — Control structures](https://go.dev/doc/effective_go#if)
