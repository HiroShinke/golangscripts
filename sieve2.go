package main

import (
	"flag"
	"fmt"
)

func primgen(n int) chan int {

	var ch chan int = make(chan int)
	go func() {
		ch <- n
		var ch1 chan int = primgen(n + 1)
		for m := range ch1 {
			if m%n == 0 {
			} else {
				ch <- m
			}
		}
	}()
	return ch
}

func main() {

	var limit *int = flag.Int("n", 30, "number of fibo numbers")
	flag.Parse()

	var ch chan int = primgen(2)

	for n := range ch {
		fmt.Print("n = ", n, "\n")
		if *limit < n {
			return
		}
	}
}
