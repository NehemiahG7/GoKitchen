package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var inv = make([][]string, 5)

func getInv() {
	fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inventory:")
	scanner := bufio.NewScanner(os.Stdin)
	fInv, err := os.Create("fInv.csv")
	if err != nil {
		log.Fatalf("File failed to create %s", err)
	}

	var group string
	var input string
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			group = "meats"
		case 1:
			group = "fruits"
		case 2:
			group = "vegetables"
		case 3:
			group = "grains"
		case 4:
			group = "dairy"
		}
		fmt.Printf("Enter any %s that you have. Seperate each by only a comma: ", group)
		scanner.Scan()
		input += group + "," + scanner.Text() + "\n"
	}
	fmt.Println(input)
	num, err := fInv.WriteString(input)
	if err != nil {
		fmt.Println(err, num)
	}
	fInv.Close()
}
func loadInv() {
	fInv, err := os.Open("fInv.csv")
	if err != nil {
		log.Fatalf("Failed to open file %s", err)
	}
	reader := bufio.NewScanner(fInv)

	for i := 0; reader.Scan(); i++ {
		str := reader.Text()
		arry := strings.Split(str, ",")
		inv[i] = arry
	}
	fInv.Close()
}
func printInv() {
	fmt.Println("Here is what you have in you kitchen:")
	for i := 0; i < len(inv); i++ {
		fmt.Printf("%s:\n\t", inv[i][0])
		for k := 1; k < len(inv[i]); k++ {
			fmt.Printf("%s, ", inv[i][k])
		}
		fmt.Printf("\n\n")
	}
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	fInv, err := os.Open("fInv.csv")

	if err != nil {
		//Initialize .csv if fInv has not been taken before
		getInv()
	}
	fInv.Close()
	loadInv()
	printInv()

}
