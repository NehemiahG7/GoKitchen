package main

import (
	"fmt"
	"os"

	"github.com/NehemiahG7/project-0/grocery"
	"github.com/NehemiahG7/project-0/index"
	"github.com/NehemiahG7/project-0/inventory"
	"github.com/NehemiahG7/project-0/util"
)

var reg string = `/.*`

func main() {

	for {

		switch {
		//Enter inventory module
		case util.CheckGegex(Module, `^\s*.?inventory\s*$`):
			Module = inventory.InventoryModule(*Inv)

		//Enter grocery module
		case util.CheckGegex(Module, `^\s*.?grocery\s*$`):
			Module = grocery.GroceryModule(*Inv, *Groc, *Index)

		//Enter index sub-module
		case util.CheckGegex(Module, `^\s*.?index\s*$`):
			Module = index.IndexModule(*Index, *Inv)

		//print all commands for main menu
		case util.CheckGegex(Module, `^\s*.?help\s*$`):
			fmt.Printf("%s\n", util.MenuHelpString)
			Module = "menu"

		//quit out of program
		case util.CheckGegex(Module, `^\s*.?q\s*$`):
			checkout()
			os.Exit(1)

		//Save on returning using checkout
		case util.CheckGegex(Module, `^\s*.?checkout\s*$`):
			checkout()
			Module = "menu"

		//menu case to prompt user for input
		case util.CheckGegex(Module, `^\s*.?menu\s*$`):
			fmt.Print("Please enter a module to load (inventory or grocery): ")
			fmt.Scanln(&Module)

		//inform of invalid command
		default:
			fmt.Printf("\nCommand, \"%s\" not found. Use, \"help\" for a list of commands\n", Module)
			Module = "menu"
		}

		// //Prompt for input
		// fmt.Print("Please enter a module to load (inventory or grocery): ")
		// fmt.Scanln(&Module)
	}

}

//Checkout encodes all currently open structs to their respective files
func checkout() {
	util.Encode(Inv, InvFile)
	util.Encode(Groc, GrocFile)
	util.Encode(Index, IdxFile)
	Groc.UpdateList(*Inv)
}
