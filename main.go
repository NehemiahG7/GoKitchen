package main

import (
	"fmt"
	"os"
)

var inv = make([][]string, 5)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	fInv, err := os.Open("fInv.json")

	if err != nil {
		//Initialize .json if fInv has not been taken before
		createInv()
	}
	fInv.Close()

	inv := loadInv()
	//inv.printInv()
	fmt.Println(inv.checkInv("Chicken"))

}
