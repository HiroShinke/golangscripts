package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	ifile := flag.String("i", "", "input file")
	ofile := flag.String("o", "", "output file")
	flag.Parse()

	fmt.Print("i = ", *ifile, "\n")
	fmt.Print("o = ", *ofile, "\n")

	data, err := os.ReadFile(*ifile)
	if err != nil {
		panic(err)
	}
	fmt.Print("data = \n", string(data))
}
