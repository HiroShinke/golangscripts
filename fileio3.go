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

	var r *bufio.Reader
	var w *bufio.Writer

	if *ifile == "" {
		r = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(*ifile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r = bufio.NewReader(f)
	}

	if *ofile == "" {
		w = bufio.NewWriter(os.Stdout)
		defer w.Flush()

	} else {
		of, _ := os.Create(*ofile)
		defer of.Close()
		w = bufio.NewWriter(of)
		defer w.Flush()
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
