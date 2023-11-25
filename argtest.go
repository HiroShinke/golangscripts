package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	for i, s := range os.Args {
		fmt.Printf("Args[%d] = %s\n", i, s)
	}

	var nFlag = flag.Int("n", 1234, "help message for flag n")
	var nString = flag.String("s", "default", "help message for flag n")
	var nCheck = flag.Bool("b", false, "boolean arg sample")
	flag.Func("c", "test Func parameter", func(arg string) error {
		fmt.Printf("c parame =%s\n", arg)
		return nil
	})

	var nFlag2 int
	var nString2 string
	var nCheck2 bool
	flag.IntVar(&nFlag2, "number", 3456, "help message for flag n")
	flag.StringVar(&nString2, "string", "default2", "help message for flag n")
	flag.BoolVar(&nCheck2, "bool", true, "help message for flag n")

	flag.Parse()
	fmt.Printf("nFlag is %T\n", nFlag)
	fmt.Print("nFlag = ", *nFlag, "\n")
	fmt.Printf("nString is %T\n", nString)
	fmt.Print("nSrring = ", *nString, "\n")
	fmt.Printf("nCheck = %t\n", *nCheck)

	fmt.Print("nFlag2 = ", nFlag2, "\n")
	fmt.Print("nSrring2 = ", nString2, "\n")
	fmt.Printf("nCheck2 = %t\n", nCheck2)

	for i, s := range flag.Args() {
		fmt.Printf("Args[%d] = %s\n", i, s)
	}

}
