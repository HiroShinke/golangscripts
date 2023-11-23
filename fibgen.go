package main

import (
	"flag"
	"fmt"
)

func create_fibgen(a int, b int) func() int {

	return func() int {
		ret := a
		a, b = b, a+b
		return ret
	}

}

func main() {

	var limit *int = flag.Int("n", 30, "number of fibo numbers")
	flag.Parse()

	gen := create_fibgen(1, 2)
	count := 0

	for {
		n := gen()
		fmt.Print("n = ", n, "\n")
		if *limit < count {
			break
		} else {
			count++
		}
	}

}
