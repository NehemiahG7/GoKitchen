package main

import (
	"os"
)

var inv = make([][]string, 5)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	fInv, err := os.Open("fInv.json")

	if err != nil {
		//Initialize .json if fInv has not been taken before
		create()
	}
	fInv.Close()

	inv := load()

	inv.editName("Beef", "Lamb")
	inv.editDate("Lamb", "Today")
	inv.editList("Lamb", true)
	inv.printInv()

}
