package inventory

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	fmt.Fprintf(&b, "\t%-16sDate added: %-16sAdd to grocery list: %t", item.Name, item.DateEntered, item.ForceList)
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
	fmt.Print(inv)
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

//Add Takes a food group and name of an item as strings and adds item to given food group with default values
func (inv Inventory) Add(strs []string) {
	date := time.Now().Format("Mon Jan 2")
	key := strs[0]
	if len(strs) == 1 {
		inv.Inven[key] = make([]Item, 0)
		return
	}

	for i := 1; i < len(strs); i++ {
		if inv.check(strs[i]) {
			inv.updateDate(strs[i])
		}
		item := Item{
			Name:        strs[i],
			ForceList:   false,
			DateEntered: date,
		}
		inv.Inven[key] = append(inv.Inven[key], item)
	}

}

//Remove takes the name of an Item and deletes it if it is in the Inventory
func (inv *Inventory) Remove(strs []string) {
	for i := 0; i < len(strs); i++ {
		key, index := inv.find(strs[i])
		if key == "" {
			continue
		}
		inv.Inven[key] = append(inv.Inven[key][:index], inv.Inven[key][index+1:]...)
	}
}

//ChangeKey moves an item to a different key value. Array should be structured in pains with the new key before each item.
func (inv Inventory) ChangeKey(arry []string) {
	//A slice to hold all items moved for the remove function later
	names := make([]string, 0)

	for i := 0; i < len(arry)/2; i++ {
		name := arry[i*2+1]
		edit := arry[i*2]
		//Find name in the Inventory, return "" if it does not currently exist
		key, index := inv.find(name)
		if key == "" {
			log.Println("Cannot find item")
			continue
		}
		//Since name exists, add it to the names array
		names = append(names, name)
		//Assign the item names refers to to a temperary variable
		tempItem := inv.Inven[key][index]
		inv.Remove(names)
		//Add the tempItem to the desired key
		inv.Inven[edit] = append(inv.Inven[edit], tempItem)
	}
}

//RemoveKey removes the given key and its []item from the Inv struct
func (inv Inventory) RemoveKey(key string) {
	delete(inv.Inven, key)
}

// editDate() changes the date of item to the edit
func (inv Inventory) updateDate(name string) {
	key, index := inv.find(name)
	date := time.Now().Format("Mon Jan 2")
	if key == "" {
		log.Println("Cannot find item")
		return
	}
	inv.Inven[key][index].DateEntered = date
}

//AddGrocery sets name's forceList to true
func (inv Inventory) AddGrocery(strs []string) {
	for i := 0; i < len(strs); i++ {
		key, index := inv.find(strs[i])
		if key == "" {
			log.Println("Cannot find item")
			return
		}
		inv.Inven[key][index].ForceList = true
	}
}

//RemoveGrocery sets name's forcelist to false
func (inv Inventory) RemoveGrocery(strs []string) {
	for i := 0; i < len(strs); i++ {
		key, index := inv.find(strs[i])
		if key == "" {
			log.Println("Cannot find item")
			return
		}
		inv.Inven[key][index].ForceList = false
	}
}

//LoadInv loads an Inventory struct from the given .json file
func LoadInv(InvFile string) *Inventory {
	file, err := os.Open(InvFile)
	if err != nil {
		//Initialize .json if fInv has not been taken before
		return createInv()
	}
	defer file.Close()
	dec := json.NewDecoder(file)

	invMap := Inventory{}
	dec.Decode(&invMap)

	return &invMap
}

//EncodeInv prints the struct to a .json file
// func (inv *Inventory) encodeInv() {
// 	encode(inv, InvFile)
// }

// create() creates a new Inventory struct from user input then calls encodeinv() to write it to a .json
func createInv() *Inventory {
	fmt.Println("It looks like you haven't used kitchen manager before. Lets take your inventory:")
	scanner := bufio.NewScanner(os.Stdin)

	str := Inventory{Inven: make(map[string][]Item)}

	//Get inventory from user
	var group string
	date := time.Now().Format("Mon Jan 2")
	for i := 0; i < 6; i++ {
		//Cycle through catagories
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
		case 5:
			group = "other"
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
	//str.encodeInv()
	return &str
}
