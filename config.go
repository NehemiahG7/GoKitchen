package main

import (
	"log"
	"os"
)

func init() {

}

func openFile(str string) *os.File {
	fInv, err := os.Open(str)
	if err != nil {
		log.Fatalf("File failed to open; %s", err)
	}
	return fInv
}
