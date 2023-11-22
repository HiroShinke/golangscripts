package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlag = flag.Int("n", 1234, "help message for flag n")
	var nString = flag.String("s", "default", "help message for flag n")
	flag.Parse()
	fmt.Print("nFlag = ", *nFlag, "\n")
	fmt.Print("nSrring = ", *nString, "\n")

}
