package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// GroceryList is a struct containing a map which uses a string for key with a slice of strings for content
type GroceryList struct {
	GList []string `json:"GroceryList"`
}

func (grocery GroceryList) String() string {
	var b strings.Builder
	for i := 0; i < len(grocery.GList); i++ {
		fmt.Fprintf(&b, "%s\n", grocery.GList[i])
	}
	return b.String()
}

//Print prints the content of a grocery struct
func (grocery GroceryList) Print() {
	fmt.Printf("\nYour grocery list:\n%s", grocery)
}

//AddToInv will add all items in the grocery list to the Inventory.
//The method will check ItemIndex to find a key to add the item at, and will add to
//'other' if a key is not found
func (grocery GroceryList) AddToInv() {
	var d strings.Builder
	fmt.Fprintf(&d, "Added: \"")
	for i := 0; i < len(grocery.GList); i++ {
		key, b := Index.CheckItem(grocery.GList[i])
		if b {
			elem := []string{key, grocery.GList[i]}
			Inv.Add(elem)
		} else {
			elem := []string{"other", grocery.GList[i]}
			Inv.Add(elem)
			fmt.Fprintf(&d, "%s, ", grocery.GList[i])
		}
	}
	fmt.Fprintf(&d, "\"to 'other'")
}

//Add appends an item to the grocery list
func (grocery GroceryList) Add(items []string) {
	for i := 0; i < len(items); i++ {
		grocery.GList = append(grocery.GList, items[i])
	}
}

//Remove removes an item from the grocery list
func (grocery GroceryList) Remove(items []string) {
	for i := 0; i < len(items); i++ {
		index := grocery.find(items[i])
		if index == -1 {
			return
		}
		grocery.GList = append(grocery.GList[:index], grocery.GList[index+1:]...)
	}
}

func (grocery GroceryList) find(item string) int {
	for i := 0; i < len(grocery.GList); i++ {
		if grocery.GList[i] == item {
			return i
		}
	}
	return -1
}

func createList(inv Inventory) *GroceryList {
	var list = make([]string, 0)
	for k := range inv.Inven {
		for i := 0; i < len(inv.Inven[k]); i++ {
			if inv.Inven[k][i].ForceList {
				list = append(list, inv.Inven[k][i].Name)
			}
		}
	}
	grocery := GroceryList{GList: list}
	return &grocery
}

func loadList() *GroceryList {
	file, err := os.Open(GrocFile)
	defer file.Close()

	if err != nil {
		//Initialize .json if fInv has not been taken before
		return createList(*Inv)
	}
	dec := json.NewDecoder(file)

	grocery := GroceryList{}
	dec.Decode(&grocery)

	return &grocery
}

// func (grocery *GroceryList) encodeList() {
// 	encode(grocery, GrocFile)
// }
