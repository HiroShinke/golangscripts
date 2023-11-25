package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {

	ifile := flag.String("i", "", "input file")
	ofile := flag.String("o", "", "output file")
	pattern := flag.String("p", "", "output file")
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

	var re, _ = regexp.Compile(*pattern)

	for {
		str, err := r.ReadString('\n')
		if err == nil || err == io.EOF {
			if re.Match([]byte(str)) {
				w.WriteString(str)
			}
		}
		if err == io.EOF {
			break
		}
	}

}
