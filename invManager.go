package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func loadInv() *Inventory {
	fInv := openFile("fInv.json")
	defer fInv.Close()
	dec := json.NewDecoder(fInv)

	invMap := Inventory{}
	dec.Decode(&invMap)
	fmt.Println(invMap)

	return &invMap
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

	fInv, err := os.Create("fInv.json")
	if err != nil {
		log.Fatalf("File failed to create %s", err)
	}

	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.Encode(inv)

	_, err = io.Copy(fInv, buf)
	if err != nil {
		fmt.Printf("Fuck: %s", err)
	}

	fInv.Close()
}

func getInv() {
	fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inventory:")
	scanner := bufio.NewScanner(os.Stdin)

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

	encodeInv(str)
}
