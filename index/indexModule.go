package index

import (
	"fmt"

	"github.com/NehemiahG7/GoKitchen/inventory"
	"github.com/NehemiahG7/GoKitchen/util"
)

var reg string = `/.*`

//IndexModule is the command loop for said module
func IndexModule(Index ItemIndex, Inv inventory.Inventory) string {
	for {
		fmt.Print("Welcome to GoKitchen - Index: ")

		commands := util.Parse(util.InputString())
		for i := 0; i < len(commands); i++ {
			switch {
			//AddItem
			case util.CheckGegex(commands[i][0], `^\s*.?add\s*$`):
				//Check if user is going to enter items with command
				if util.CheckGegex(commands[i][0], reg) {
					Index.AddItem(commands[i][1], commands[i][2])
					fmt.Println("In Quick Add")
					continue
				}
				//request items
				fmt.Println("What would you like to add?\nPlease seperate all items with a coma:")
				arry := util.ParseLine(util.InputString())
				Index.AddItem(arry[0], arry[1])

			//RemoveItem
			case util.CheckGegex(commands[i][0], `^\s*.?remove\s*$`):
				//Check if user is going to enter items with command
				if util.CheckGegex(commands[i][0], reg) {
					Index.RemoveItem(commands[i][1])
					fmt.Println("In Quick Add")
					continue
				}
				//request items
				fmt.Println("What would you like to remove?\nPlease seperate all items with a coma:")
				arry := util.ParseLine(util.InputString())
				Index.RemoveItem(arry[0])

			//PrintIndex
			case util.CheckGegex(commands[i][0], `^\s*.?print\s*$`):
				Index.PrintIndex()

			//Help
			case util.CheckGegex(commands[i][0], `^\s*.?help\s*$`):
				fmt.Println(util.IndexHelpString)

			case util.CheckGegex(commands[i][0], `^\s*.?exit\s*$`):
				return "checkout"

			case util.CheckGegex(commands[i][0], `^\s*.?q\s*$`):
				return "q"

			default:
				fmt.Printf("Command, \"%s\" not found, continuing with next command", commands[i][0])
			}
		}
	}
}
