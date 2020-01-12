package main

import (
	"log"
	"os"
)

func openFile() *os.File {
	fInv, err := os.Open("fInv.csv")
	if err != nil {
		log.Fatalf("File failed to open; %s", err)
	}
	return fInv
}
