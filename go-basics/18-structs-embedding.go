// Effective Go — structs
// https://go.dev/doc/effective_go
//
package gobasics

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Dept string
}

func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func (p Person) Greeting() string {
	return "Hi, " + p.Name
}

// User — ví dụ slice of struct trong API/DB layer.
type User struct {
	ID   int
	Name string
}

// UsersByValue — slice of struct (value).
func UsersByValue() []User {
	return []User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
	}
}

// UsersByPointer — slice of pointers; thường dùng khi struct lớn hoặc chia sẻ instance.
func UsersByPointer() []*User {
	return []*User{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
	}
}

// EmptyUserPointers — make([]*User, 0) vs nil: empty, không nil, sẵn sàng append.
func EmptyUserPointers() []*User {
	return make([]*User, 0)
}

// AppendUserValue appends a User value (struct copied into slice).
func AppendUserValue(users []User, u User) []User {
	return append(users, u)
}

// AppendUserPointer appends *User (chỉ copy pointer vào slice).
func AppendUserPointer(users []*User, u *User) []*User {
	return append(users, u)
}

// RenameUserAt đổi tên user tại index — []User sửa trực tiếp phần tử trong slice.
func RenameUserAt(users []User, index int, name string) {
	if index < 0 || index >= len(users) {
		return
	}
	users[index].Name = name
}

// RenameUserPtr đổi qua *User — mọi slice/code cùng trỏ tới pointer đều thấy đổi.
func RenameUserPtr(u *User, name string) {
	if u == nil {
		return
	}
	u.Name = name
}

// SharedPointerDemo: hai slice entries cùng trỏ một *User.
func SharedPointerDemo() (first, second *User, sameInstance bool) {
	u := &User{ID: 1, Name: "John"}
	slice := []*User{u, u}
	return slice[0], slice[1], slice[0] == slice[1]
}
