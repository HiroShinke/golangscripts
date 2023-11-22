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

	if *ofile == "" {
		fmt.Print(string(data))
	} else {
		err := os.WriteFile(*ofile, data, 0644)
		if err != nil {
			panic(err)
		}
	}
}
