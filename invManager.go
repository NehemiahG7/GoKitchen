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
	"time"
)

// Inventory is a struct containing a map of Items, using food groups as keys
type Inventory struct {
	Inven map[string][]Item `json:"Inventory"`
}

// Item is a struct containing a string, Name, bool, ForceList, and string, DateEntered
type Item struct {
	Name        string `json:"name"`
	ForceList   bool   `json:"forceList"`
	DateEntered string `json:"dateEntered"`
}

// String() returns string version of Item
func (item Item) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "\t%-16sDate added- %-16sAlways add to grocery list: %t", item.Name, item.DateEntered, item.ForceList)
	return b.String()
}

// String() returns string version of Inventory
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

// printInv() prints Inventory
func (inv Inventory) printInv() {
	fmt.Println("Here is what you have in you kitchen:")
	fmt.Println(inv)
}

// check() takes the name of an item and returns true if it is in Inventory, and false if not
func (inv Inventory) check(item string) bool {
	contains := false
	for k := range inv.Inven {
		for i := 0; i < len(inv.Inven[k]); i++ {
			if inv.Inven[k][i].Name == item {
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

// find takes a string and returns the key and index of the corrosponding item in the Inventory map.
// The function returns two nil values if the item is not present.
func (inv Inventory) find(name string) (string, int) {
	var str string
	var num int
	for k := range inv.Inven {
		for i := 0; i < len(inv.Inven[k]); i++ {
			if inv.Inven[k][i].Name == name {
				return k, i
			}
		}
	}
	return str, num
}

// add() Takes a food group and name of an item as strings and adds item to given food group with default values
func (inv Inventory) add(key string, name string) {
	date := time.Now().Format("Mon Jan 2")

	item := Item{
		Name:        name,
		ForceList:   false,
		DateEntered: date,
	}
	inv.Inven[key] = append(inv.Inven[key], item)
}

// deleteItem() takes the name of an Item and deletes it if it is in the Inventory
func (inv Inventory) deleteItem(name string) {
	key, index := inv.find(name)
	if key == "" {
		return
	}
	inv.Inven[key] = append(inv.Inven[key][:index], inv.Inven[key][index+1:]...)
}

// editName() changes the name of name to edit
func (inv Inventory) editName(name string, edit string) {
	key, index := inv.find(name)
	if key == "" {
		log.Println("Cannot find item")
		return
	}
	inv.Inven[key][index].Name = edit
}

// editDate() changes the date of item to the edit
func (inv Inventory) editDate(name string, edit string) {
	key, index := inv.find(name)
	if key == "" {
		log.Println("Cannot find item")
		return
	}
	inv.Inven[key][index].DateEntered = edit
}

// editList() sets name's forceList to edit
func (inv Inventory) editList(name string, edit bool) {
	key, index := inv.find(name)
	if key == "" {
		log.Println("Cannot find item")
		return
	}
	inv.Inven[key][index].ForceList = edit
}

// load() loads an Inventory struct from the given .json file
func load() *Inventory {
	fInv := openFile("fInv.json")
	defer fInv.Close()
	dec := json.NewDecoder(fInv)

	invMap := Inventory{}
	dec.Decode(&invMap)

	return &invMap
}

// encodeInv() takes an Inventory struct, encodes it to json, and writes it to a .json file
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
		fmt.Printf("%s", err)
	}

	fInv.Close()
}

// create() creates a new Inventory struct from user input then calls encodeinv() to write it to a .json
func create() {
	fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inventory:")
	scanner := bufio.NewScanner(os.Stdin)

	str := Inventory{Inven: make(map[string][]Item)}

	//Get inventory from user
	var group string
	date := time.Now().Format("Mon Jan 2")
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
			arry[k] = Item{Name: itmArray[k], DateEntered: date, ForceList: false}
		}

		//add group to inventory
		str.Inven[group] = arry
	}

	encodeInv(str)
}
