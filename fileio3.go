package main

import (
	"bufio"
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
	r := bufio.NewReader(f)

	var w *bufio.Writer

	if *ofile == "" {
		w = bufio.NewWriter(os.Stdout)
		defer w.Flush()

	} else {
		of, err = os.Create(*ofile)
		defer of.Close()
		w = bufio.NewWriter(of)
	}

	buff := make([]byte, 512)
	for {
		n, err := r.Read(buff)
		if err == nil || err == io.EOF {
			w.Write(buff[:n])
		}
		if err == io.EOF {
			break
		}
	}

}
