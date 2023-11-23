package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	ifile := flag.String("i", "", "input file")
	ofile := flag.String("o", "", "output file")
	flag.Parse()

	fmt.Print("i = ", *ifile, "\n")
	fmt.Print("o = ", *ofile, "\n")

	f, err := os.Open(*ifile)
	if err != nil {
		panic(err)
	}

	var of *os.File

	if *ofile == "" {
		of = os.Stdout

	} else {
		of, err = os.Create(*ofile)
		defer of.Close()
	}

	buff := make([]byte, 512)
	for {
		n, err := f.Read(buff)
		if err == nil || err == io.EOF {
			of.Write(buff[:n])
		}
		if err == io.EOF {
			break
		}
	}

}
