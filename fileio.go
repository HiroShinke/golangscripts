package main

import (
	"flag"
	"fmt"
)

func main() {

	ifile := flag.String("i", "", "input file")
	ofile := flag.String("o", "", "output file")
	flag.Parse()

	fmt.Print("i = ", *ifile)
	fmt.Print("o = ", *ofile)
}
