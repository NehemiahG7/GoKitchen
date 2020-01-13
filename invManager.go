package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type InvMethods interface {
	loadInv()
	printInv()
}
type Inventory struct {
	Inven map[string]*[]Item `json:"Inventory"`
}
type Item struct {
	Name        string `json:"name"`
	ForceList   bool   `json:"forceList"`
	DateEntered string `json:"dateEntered"`
}

func loadInv() {
	fInv := openFile()
	reader := bufio.NewScanner(fInv)

	//.Scan() through lines of fInv.csv
	for i := 0; reader.Scan(); i++ {
		//Add all lines to the inv array
		inv[i] = strings.Split(reader.Text(), ",")
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

func encodeInv(inv Inventory) {
	fInv := openFile()

	encodeInv, _ := json.MarshalIndent(inv, "", "  ")
	fmt.Println(encodeInv)
	//export
	num, err := fInv.Write(encodeInv)
	if err != nil {
		fmt.Println(err, num)
		//get
	}
	fInv.Close()
}

func getInv() {
	fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inventory:")
	scanner := bufio.NewScanner(os.Stdin)

	//Make file for later use
	fInv, err := os.Create("fInv.json")
	if err != nil {
		log.Fatalf("File failed to create %s", err)
	}
	fInv.Close()

	str := Inventory{Inven: make(map[string]*[]Item)}

	//Get inventory from user
	var group string
	for i := 0; i < 5; i++ {
		//Cycle through catagories
		switch i {
		case 0:
			group = "Meats"
		case 1:
			group = "Fruits"
		case 2:
			group = "Vegetables"
		case 3:
			group = "Grains"
		case 4:
			group = "Dairy"
		}
		fmt.Printf("Enter any %s that you have. Seperate each by only a comma: ", group)
		scanner.Scan()

		//Split input into array
		itmArray := strings.Split(scanner.Text(), ",")

		//Make array to group items together
		arry := make([]Item, len(itmArray))

		//make structs for each item
		for k := 0; k < len(itmArray); k++ {
			arry[k] = Item{Name: itmArray[k], DateEntered: "p", ForceList: false}
		}

		//add group to inventory
		str.Inven[group] = &arry
	}

	fmt.Println(str)

	encodeInv(str)
}
