package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    var fn0 = 0
	var fn1 = 1
	var fn = 0
	var n = 0

	return func() int {
		if n == 0 {
			fn = fn0
		} else if n == 1 {
			fn = fn1
		} else if n > 1 {
			fn = fn1 + fn0
			fn0 = fn1
			fn1 = fn
		}
		n++

		return fn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
