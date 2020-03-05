package grocery

import (
	"fmt"

	"github.com/NehemiahG7/project-0/index"
	"github.com/NehemiahG7/project-0/inventory"
	"github.com/NehemiahG7/project-0/util"
)

var reg string = `/.*`

func GroceryModule(inv inventory.Inventory, groc GroceryList, idx index.ItemIndex) string {
	for {
		fmt.Print("Welcome to GoKitchen - Grocery: ")

		commands := util.Parse(util.InputString())
		for i := 0; i < len(commands); i++ {
			switch {
			//Add
			case util.CheckGegex(commands[i][0], `^\s*.?add\s*$`):
				//Check if user is going to enter items with command
				if util.CheckGegex(commands[i][0], reg) {
					groc.Add(commands[i][1:])
					continue
				}
				//request items
				fmt.Println("What would you like to add?\nPlease seperate all items with a coma:")
				groc.Add(util.ParseLine(util.InputString()))

			//Remove
			case util.CheckGegex(commands[i][0], `^\s*.?remove\s*$`):
				//Check if user is going to enter items with command
				if util.CheckGegex(commands[i][0], reg) {
					groc.Remove(commands[i][1:])
					continue
				}
				//request items
				fmt.Println("What would you like to remove?\nPlease seperate all items with a coma:")
				groc.Remove(util.ParseLine(util.InputString()))

			//AddToInv
			case util.CheckGegex(commands[i][0], `^\s*.?addToInv\s*$`):
				groc.AddToInv(inv, idx)
				fmt.Println("Inventory updated from grocery list")

			//ExportGrocery
			case util.CheckGegex(commands[i][0], `^\s*.?export\s*$`):
				groc.ExportList()
				fmt.Println("Grocery list saved to groceryList.txt")

			//Print the grocery list
			case util.CheckGegex(commands[i][0], `^\s*.?print\s*$`):
				groc.Print()

			//help
			case util.CheckGegex(commands[i][0], `^\s*.?help\s*$`):
				fmt.Println(util.GroceryHelpString)

			//move to index sub-module
			case util.CheckGegex(commands[i][0], `^\s*.?index\s*$`):
				//index.IndexModule(idx, inv)
				return ""

			//return to main
			case util.CheckGegex(commands[i][0], `^\s*.?exit\s*$`):
				return "checkout"

			//quit program
			case util.CheckGegex(commands[i][0], `^\s*.?q\s*$`):
				return "q"

			//invalid command
			default:
				fmt.Printf("Command, \"%s\" not found, continuing with next command", commands[i][0])
			}
		}
	}
}
