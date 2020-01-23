package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	//string builder to report all items added with key "other"
	var d strings.Builder
	fmt.Fprintf(&d, "Added: \"")

	//Iterate through grocery list. Checking the Item index for a catagory for each item
	//Assign "Other" if a key does not exist
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
func (grocery *GroceryList) Add(items []string) {
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

//finds an item and returns its index. Returns -1 if item is not in list
func (grocery GroceryList) find(item string) int {
	for i := 0; i < len(grocery.GList); i++ {
		if grocery.GList[i] == item {
			return i
		}
	}
	return -1
}

//UpdateList updates the provided grocery list from the global inventory item
func (grocery *GroceryList) UpdateList() {
	var arry = make([]string, 0)
	for k := range Inv.Inven {
		for i := 0; i < len(Inv.Inven[k]); i++ {
			if Inv.Inven[k][i].ForceList {
				if grocery.find(Inv.Inven[k][i].Name) != -1 {
					arry = append(arry, Inv.Inven[k][i].Name)
				}
			}
		}
	}
	grocery.Add(arry)
}

//ExportList exports the used grocery list to a text file
func (grocery GroceryList) ExportList() {
	list, err := os.Create("groceryList.txt")
	defer list.Close()
	if err != nil {
		log.Fatalf("Error, could not open groceryList.txt, %s", err)
	}
	_, err = list.WriteString(grocery.String())
	if err != nil {
		log.Fatalf("Error, could not write to groceryList.txt, %s", err)
	}
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

//load list from grocery file and return its pointer
//call createList if file does not exist
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
