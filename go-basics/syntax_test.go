package gobasics

import "testing"

func TestFizzBuzz(t *testing.T) {
	if FizzBuzz(15) != "FizzBuzz" || FizzBuzz(3) != "Fizz" || FizzBuzz(7) != "7" {
		t.Fatal(FizzBuzz(15), FizzBuzz(3), FizzBuzz(7))
	}
}

func TestReverseInPlace(t *testing.T) {
	s := []int{1, 2, 3, 4}
	ReverseInPlace(s)
	want := []int{4, 3, 2, 1}
	for i := range want {
		if s[i] != want[i] {
			t.Fatalf("%v", s)
		}
	}
}

func TestWordCount(t *testing.T) {
	m := WordCount("go is go")
	if m["go"] != 2 || m["is"] != 1 {
		t.Fatalf("%v", m)
	}
}

func TestCounterPointerReceiver(t *testing.T) {
	var c Counter
	c.Inc()
	c.Add(4)
	if c.Value() != 5 {
		t.Fatal(c.Value())
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Fatal(a, b)
	}
}

func TestDeferOrder(t *testing.T) {
	got := DeferOrder()
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Fatalf("%v", got)
	}
}

func TestDeferSum(t *testing.T) {
	if DeferSum([]int{1, 2, 3}) != 12 { // (1+2+3)*2
		t.Fatal()
	}
}

func TestExportedName(t *testing.T) {
	if ExportedName() != "exported" {
		t.Fatal()
	}
}
