package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inv, err := os.Create("inv.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	fmt.Println("Enter what you have in your pantry seperated by comas: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(input)
	arry := strings.Split(input, ",")

	for i := 0; i < len(arry); i++ {
		num, err := inv.WriteString(arry[i])
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		fmt.Println(num)
	}
	inv.Close()

}
