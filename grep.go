package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func do_rec_file(
	path string,
	proc func(string, fs.FileInfo) error) error {

	info, err := os.Stat(path)
	err = proc(path, info)
	if err != nil {
		return err
	}

	if info.IsDir() && info.Name() != ".git" {
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			path := filepath.Join(path, entry.Name())
			if err := do_rec_file(path, proc); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}

func main() {

	ifile := flag.String("i", "", "input file")
	pattern := flag.String("pattern", "", "input file")
	flag.Parse()
	fmt.Print("i = ", *ifile, "\n")
	fmt.Print("pattern = ", *pattern, "\n")

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	proc := func(path string, info fs.FileInfo) error {

		if !info.IsDir() {

			re, err := regexp.Compile(*pattern)
			if err != nil {
				panic(err)
			}
			f, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			r := bufio.NewReader(f)

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

		return nil
	}

	if err := do_rec_file(*ifile, proc); err != nil {
		fmt.Println(err)
	}

}
