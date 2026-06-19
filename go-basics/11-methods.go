// Methods — value vs pointer receiver
//
// func (p Person) M() {}     // copy
// func (p *Person) M() {}    // can mutate p
// use pointer receiver when method mutates or struct is large
package gobasics

type Counter struct {
	n int
}

func (c Counter) Value() int {
	return c.n
}

func (c *Counter) Inc() {
	c.n++
}

func (c *Counter) Add(delta int) {
	c.n += delta
}
