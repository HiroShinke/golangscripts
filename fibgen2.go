package main

import (
	"flag"
	"fmt"
)

func main() {

	var limit *int = flag.Int("n", 30, "number of fibo numbers")
	flag.Parse()

	var ch chan int = make(chan int)

	var create_fibgen = func(a int, b int) func() {
		return func() {
			ret := a
			a, b = b, a+b
			ch <- ret
		}
	}

	go func() {

		gen := create_fibgen(1, 2)
		count := 0

		for {
			gen()
			if *limit < count {
				break
			} else {
				count++
			}
		}
		close(ch)

	}()

	for n := range ch {
		fmt.Print("n = ", n, "\n")
	}

}
