package index

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/NehemiahG7/project-0/inventory"
)

//ItemIndex is a map containing keys that are all item names ever entered into Inv, matched with the catagory they were entered for
type ItemIndex struct {
	Idx map[string]string `json:"Index"`
}

func (index ItemIndex) String() string {
	var b strings.Builder
	for i := range index.Idx {
		fmt.Fprintf(&b, "%-10s: %s\n", i, index.Idx[i])
	}
	return b.String()
}

//PrintIndex prints the provided index
func (index ItemIndex) PrintIndex() {
	fmt.Println(index)
}

//AddItem adds an item to the index, or updates it's catagory if the item already exists
func (index ItemIndex) AddItem(name string, key string) {
	index.Idx[name] = key
}

//RemoveItem will item from the index
func (index ItemIndex) RemoveItem(item string) {
	delete(index.Idx, item)
}

//CheckItem checks if item is a key in the index and returns it's element and a boolean
func (index ItemIndex) CheckItem(item string) (string, bool) {
	elem, b := index.Idx[item]
	return elem, b
}

// //ExportIndex saves the index to a .json file
// func (index ItemIndex) ExportIndex() {
// 	encode(index, IndexFile)
// }

func (index ItemIndex) updateIndex(inv inventory.Inventory) {
	for k := range inv.Inven {
		if k == "other" {
			continue
		}
		for i := 0; i < len(inv.Inven[k]); i++ {
			_, b := index.CheckItem(inv.Inven[k][i].Name)
			if !b {
				index.AddItem(inv.Inven[k][i].Name, k)
			}
		}
	}
}

//LoadIndex loads an ItemIndex from the file name given in conf.json. If the file does
//not exist, createIndex is called
func LoadIndex(inv inventory.Inventory, indexFile string) *ItemIndex {
	file, err := os.Open(indexFile)
	defer file.Close()
	if err != nil {
		//Initialize .json if fInv has not been taken before
		return createIndex(inv)
	}

	dec := json.NewDecoder(file)

	index := ItemIndex{}
	dec.Decode(&index)
	index.updateIndex(inv)
	return &index
}

//create and item index from the given inv
func createIndex(inv inventory.Inventory) *ItemIndex {
	index := ItemIndex{}
	mp := make(map[string]string)
	for k := range inv.Inven {
		if k == "other" {
			continue
		}
		for i := 0; i < len(inv.Inven[k]); i++ {
			mp[inv.Inven[k][i].Name] = k
		}
	}
	index.Idx = mp
	return &index
}
