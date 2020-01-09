package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInv() {
	scanner := bufio.NewScanner(os.Stdin)
	inv, err := os.Create("inv.csv")
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
	num, err := inv.WriteString(input)
	if err != nil {
		fmt.Println(err, num)
	}
	inv.Close()
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	inv, err := os.Open("inv.csv")

	if err != nil {
		//Initialize .csv if inv has not been taken before
		fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inv:")
		getInv()
	}
	inv.Close()

}
