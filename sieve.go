package main

import (
	"flag"
	"fmt"
)

func primgen(ch *(chan int), n int) {

	*ch <- n

	var ch1 chan int = make(chan int)
	go primgen(&ch1, n+1)

	for m := range ch1 {
		if m%n == 0 {

		} else {
			*ch <- m
		}
	}
}

func main() {

	var limit *int = flag.Int("n", 30, "number of fibo numbers")
	flag.Parse()

	var ch chan int = make(chan int)

	go primgen(&ch, 2)

	for n := range ch {
		fmt.Print("n = ", n, "\n")
		if *limit < n {
			return
		}
	}
}
