// Structs — fields, literals, anonymous embedding, tags (json:"name")
//
// type User struct { Name string; Age int }
// u := User{Name: "Ada", Age: 30}
// anonymous: type Admin struct { User; Level int }
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
