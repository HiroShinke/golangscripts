package main

import "fmt"

func main() {

	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, i := range digits {
		for _, j := range digits {
			fmt.Print(i*j, " ")
		}
		fmt.Print("\n")
	}
}
