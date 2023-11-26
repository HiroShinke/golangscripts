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
	if err != nil {
		fmt.Print(err)
		return err
	}
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
	excext := flag.String("excext", "", "input file")
	incext := flag.String("incext", "", "input file")
	flag.Parse()
	fmt.Print("i = ", *ifile, "\n")
	fmt.Print("pattern = ", *pattern, "\n")

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var exccond func(string) bool
	if *excext != "" {
		re, err := regexp.Compile(*excext)
		if err != nil {
			panic(err)
		}
		exccond = func(name string) bool {
			return !re.Match([]byte(name))
		}
	} else {
		exccond = func(name string) bool {
			return true
		}
	}

	var inccond func(string) bool
	if *incext != "" {
		re, err := regexp.Compile(*incext)
		if err != nil {
			panic(err)
		}
		inccond = func(name string) bool {
			return re.Match([]byte(name))
		}
	} else {
		inccond = func(name string) bool {
			return true
		}
	}

	var re *regexp.Regexp
	var err error
	if *pattern != "" {
		re, err = regexp.Compile(*pattern)
		if err != nil {
			panic(err)
		}
	} else {
		panic("pattren is required!!")
	}

	proc := func(path string, info fs.FileInfo) error {

		if !info.IsDir() && exccond(info.Name()) && inccond(info.Name()) {

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
