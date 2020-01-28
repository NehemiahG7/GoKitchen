package main

import (
	"fmt"
	"os"

	_ "github.com/NehemiahG7/project-0/grocery"
	_ "github.com/NehemiahG7/project-0/index"
	_ "github.com/NehemiahG7/project-0/inventory"
	_ "github.com/NehemiahG7/project-0/util"
)

var reg string = `/.*`

func main() {

	for {

		switch {
		//Enter inventory module
		case CheckGegex(Module, `^\s*.?inventory\s*$`):
			InventoryModule(Inv)

		//Enter grocery module
		case CheckGegex(Module, `^\s*.?grocery\s*$`):
			GroceryModule(Inv, Groc, Index)

		//Enter index sub-module
		case CheckGegex(Module, `^\s*.?index\s*$`):
			IndexModule(Inv)

		//print all commands for main menu
		case CheckGegex(Module, `^\s*.?help\s*$`):
			fmt.Printf("Comands:\n\t-inventory: loads the inventory module\n\t-grocery: loads the grocery module\n\t-q: exits the program")

		//quit out of program
		case CheckGegex(Module, `^\s*.?q\s*$`):
			Checkout()
			os.Exit(1)

		//empty case to prevent default case from running when not using a flag
		case CheckGegex(Module, `^\s*.?empty\s*$`):

		//inform of invalid command
		default:
			fmt.Printf("\nCommand, \"%s\" not found. Use, \"help\" for a list of commands\n", Module)
		}

		//Prompt for input
		fmt.Print("Please enter a module to load (inventory or grocery): ")
		fmt.Scanln(&Module)
	}

}

//Checkout encodes all currently open structs to their respective files
func Checkout() {
	Encode(Inv, InvFile)
	Encode(Groc, GrocFile)
	Encode(Index, IndexFile)
	Groc.UpdateList()
}
