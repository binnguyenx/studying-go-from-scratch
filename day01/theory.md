# Day 1 — Cập nhật lý thuyết (Setup + Basics)

> Chỉ tóm tắt kiến thức, không làm bài tập ở đây. Code tham khảo: `day01.go`.

## Mục tiêu

- Cài Go, editor (VS Code + extension Go).
- Hiểu cách khai báo biến / hằng và các kiểu cơ bản: `string`, `int`, `float64`, `bool`.
- In ra terminal với `fmt`.

## Cài Go

- Tải bản phù hợp OS: [go.dev/dl](https://go.dev/dl/)
- Kiểm tra: `go version` — nên dùng Go 1.21 trở lên.
- Hằng ngày làm việc trong thư mục project có `go.mod`; không cần nhớ sâu `GOPATH` trừ khi gặp lỗi cấu hình.

## VS Code + extension Go

- VS Code: [code.visualstudio.com](https://code.visualstudio.com/)
- Extension **Go** (Go Team at Google): [marketplace](https://marketplace.visualstudio.com/items?itemName=golang.go)
- Giúp format, nhảy định nghĩa, gợi ý chạy / test.

## Chương trình tối thiểu

- File chạy được: `package main` + `func main()`.
- Dùng thư viện chuẩn: `import "fmt"` rồi gọi `Println` / `Printf` trong `main`.

## Hằng (`const`)

- Giá trị cố định lúc **biên dịch**, không gán lại được.
- Có thể khai báo từng dòng hoặc khối `const ( ... )`.
- `iota`: bộ đếm **trong một khối `const`**, tăng dần theo dòng (thường dùng cho enum kiểu số).

## Biến — `var` và `:=`

| Cách | Ví dụ | Ghi chú |
|------|--------|--------|
| `var` có kiểu | `var name string` | Rõ kiểu; có thể gán sau; dùng ở cấp package hoặc khi cần zero value rõ ràng |
| `var` + giá trị | `var year int = 2009` | Có thể bỏ kiểu nếu compiler suy ra được |
| `:=` | `msg := "hi"` | **Chỉ trong hàm**; khai báo + gán; kiểu suy luận từ vế phải |
| Khối `var` | `var ( a = 1; b string )` | Gom nhiều biến cho dễ đọc |

- Khai báo song song: `a, b, c := 1, 2.5, true` — mỗi vế tương ứng một biến.

## Giá trị zero

Nếu khai báo mà **chưa gán**, Go gán giá trị mặc định:

- Số: `0`
- `string`: `""`
- `bool`: `false`
- Con trỏ, slice, map, channel, interface, function: `nil`

## Kiểu cơ bản (ngày 1)

- **`string`**: chuỗi UTF-8; `len(s)` là **số byte**, không phải số ký tự Unicode trực tiếp.
- **`int`**: số nguyên; còn có `int8` … `int64`, `uint` … — mặc định literal số nguyên thường là `int`.
- **`float64`**: số thực; literal như `3.14` thường là `float64`.
- **`bool`**: chỉ `true` hoặc `false`.

## Ép kiểu

- Go **không** tự ép kiểu số hỗn hợp trong biểu thức; phải viết rõ, ví dụ: `float64(n)`, `int(x)`.

## In — gói `fmt`

- **`Println`**: in nhiều giá trị, tự thêm khoảng trắng giữa các đối số, xuống dòng cuối.
- **`Printf`**: chuỗi định dạng + đối số; “động từ” thường gặp: `%s` (chuỗi), `%d` (số nguyên thập phân), `%f` / `%.2f` (thực), `%t` (bool), `%v` (mặc định), `%T` (kiểu), `%%` (dấu `%`).
- **`Sprintf`**: giống `Printf` nhưng **trả về `string`**, không in ra stdout.

Tài liệu: [pkg.go.dev/fmt](https://pkg.go.dev/fmt).

## Con trỏ (nhắc sớm)

- `&x`: địa chỉ của `x`; kiểu `*T` là “con trỏ tới T”.
- Giá trị zero của con trỏ là `nil`; dereference `*p` chỉ an toàn khi `p` trỏ tới ô nhớ hợp lệ.

## Đọc thêm

- [Tour of Go](https://go.dev/tour/list) — mục Packages, Variables, Functions.
- [Go Proverbs](https://go-proverbs.github.io/)
- [50 Shades of Go](https://golang50shades.com/) — đọc lướt các lỗi thường gặp.
