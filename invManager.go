package main

import (
	"bufio"
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
	Inven []FoodGroup `json:"Inventory"`
}
type FoodGroup struct {
	Group string `json:"FoodGroup"`
	Items []Item `json:"items"`
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

func getInv() {
	fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inventory:")
	scanner := bufio.NewScanner(os.Stdin)
	fInv, err := os.Create("fInv.json")
	if err != nil {
		log.Fatalf("File failed to create %s", err)
	}

	str := Inventory{Inven: make([]FoodGroup, 5)}

	//Get inventory from user
	var group string
	var input string
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

		//Make struct for food group
		fdGroup := FoodGroup{Group: group, Items: make([]Item, len(itmArray))}

		//make structs for each item
		for k := 0; k < len(itmArray); k++ {
			fdGroup.Items[k] = Item{Name: itmArray[k], DateEntered: "p", ForceList: false}
		}

		//add group to inventory
		str.Inven[i] = fdGroup
	}

	fmt.Println(str)

	//export
	num, err := fInv.WriteString(input)
	if err != nil {
		fmt.Println(err, num)
		//get
	}
	fInv.Close()
}
