// Effective Go — methods.go
// https://go.dev/doc/effective_go
//
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
