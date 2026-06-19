// if / else if / else — switch (with or without tag)
//
// if x := f(); x > 0 { }  // init statement scope
// switch x { case 1,2: default: }
// switch { case x < 0: }   // tagless switch
// no automatic fallthrough (unless fallthrough)
package gobasics

func Sign(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	return 1
}

func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return itoa(n)
	}
}

func WeekdayName(d int) string {
	switch d {
	case 1:
		return "Mon"
	case 2:
		return "Tue"
	case 3:
		return "Wed"
	case 4:
		return "Thu"
	case 5:
		return "Fri"
	default:
		return "Weekend"
	}
}

// itoa tiny helper to avoid strconv import in syntax demo.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	if neg {
		return "-" + string(digits)
	}
	return string(digits)
}
