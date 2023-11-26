package main

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"log"
)

func main() {
	// 髙(FBFC)
	s, _, err := transform.String(japanese.ShiftJIS.NewDecoder(), "\xFB\xFC")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(s)
	// 﨑(FAB1)
	s, _, err = transform.String(japanese.ShiftJIS.NewDecoder(), "\xFA\xB1")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(s)
}
