package main

import (
	"os"
)

var inv = make([][]string, 5)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	fInv, err := os.Open("fInv.json")

	if err != nil {
		//Initialize .csv if fInv has not been taken before
		getInv()
	}
	fInv.Close()
	loadInv()
	printInv()

}
