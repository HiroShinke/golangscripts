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
	flag.Parse()
	fmt.Print("i = ", *ifile, "\n")
	fmt.Print("pattern = ", *pattern, "\n")

	var ch chan string = make(chan string)
	var worker chan int = make(chan int)
	re, err := regexp.Compile(*pattern)
	if err != nil {
		panic(err)
	}
	var count int = 0

	proc := func(path string, info fs.FileInfo) error {

		if !info.IsDir() {

			count++

			go func() {

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
							ch <- str
						}
					}
					if err == io.EOF {
						break
					}
				}
				worker <- 1
			}()
		}

		return nil
	}

	if err := do_rec_file(*ifile, proc); err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

loop:
	for {
		select {
		case str, _ := <-ch:
			if str != "" {
				w.WriteString(str)
			} else {
				break loop
			}
		case <-worker:
			count--
			if count == 0 {
				close(worker)
				close(ch)
				worker = nil
			}
		}
	}
}
