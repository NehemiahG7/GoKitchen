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
	Inven map[string][]Item `json:"Inventory"`
}
type Item struct {
	Name        string `json:"name"`
	ForceList   bool   `json:"forceList"`
	DateEntered string `json:"dateEntered"`
}

func (item Item) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "\t%-16sDate added- %-10sAlways add to grocery list: %t", item.Name, item.DateEntered, item.ForceList)
	return b.String()
}
func (inv Inventory) String() string {
	var b strings.Builder
	for k := range inv.Inven {
		fmt.Fprintf(&b, "%s:\n", k)
		for i := 0; i < len(inv.Inven[k]); i++ {
			fmt.Fprintf(&b, "%s\n", inv.Inven[k][i])
		}
	}
	return b.String()
}

func (inv Inventory) printInv() {
	fmt.Println("Here is what you have in you kitchen:")
	fmt.Println(inv)
}
func loadInv() *Inventory {
	fInv := openFile("fInv.json")
	defer fInv.Close()
	dec := json.NewDecoder(fInv)

	invMap := Inventory{}
	dec.Decode(&invMap)

	return &invMap
}

func (inv Inventory) checkInv(str string) bool {
	contains := false
	for k := range inv.Inven {
		for i := 0; i < len(inv.Inven[k]); i++ {
			if inv.Inven[k][i].Name == str {
				contains = true
				break
			}
		}
		if contains {
			break
		}
	}
	return contains
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

func createInv() {
	fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inventory:")
	scanner := bufio.NewScanner(os.Stdin)

	str := Inventory{Inven: make(map[string][]Item)}

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
		str.Inven[group] = arry
	}

	encodeInv(str)
}
