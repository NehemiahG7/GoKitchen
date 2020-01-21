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
	fmt.Print(grocery)
}

//Add appends an item to the grocery list
func (grocery GroceryList) Add(item string) {
	grocery.GList = append(grocery.GList, item)
}

//Remove removes an item from the grocery list
func (grocery GroceryList) Remove(item string) {
	index := grocery.find(item)
	if index == -1 {
		return
	}
	grocery.GList = append(grocery.GList[:index], grocery.GList[index+1:]...)
}

func (grocery GroceryList) find(item string) int {
	for i := 0; i < len(grocery.GList); i++ {
		if grocery.GList[i] == item {
			return i
		}
	}
	return -1
}

func createList(inv Inventory) GroceryList {
	var list = make([]string, 10)
	for k := range inv.Inven {
		for i := 0; i < len(inv.Inven[k]); i++ {
			if inv.Inven[k][i].ForceList {
				list = append(list, inv.Inven[k][i].Name)
			}
		}
	}
	grocery := GroceryList{GList: list}
	return grocery
}

func loadList() *GroceryList {
	file, err := os.Open(GrocFile)
	if err != nil {
		//Initialize .json if fInv has not been taken before
		//return createList()
		fmt.Printf("Could not open file, %s", err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)

	grocery := GroceryList{}
	dec.Decode(&grocery)

	return &grocery
}

func encodeList(grocery GroceryList) {
	encode(grocery, GrocFile)
}
