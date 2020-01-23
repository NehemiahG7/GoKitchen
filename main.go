package main

import (
	"fmt"
	"os"
)

var reg string = `/.*`

func main() {

	for {

		switch {
		//Enter inventory module
		case checkGegex(Module, `^\s*.?inventory\s*$`):
			inventoryModule()

		//Enter grocery module
		case checkGegex(Module, `^\s*.?grocery\s*$`):
			groceryModule()

		//Enter index sub-module
		case checkGegex(Module, `^\s*.?index\s*$`):
			indexModule()

		//print all commands for main menu
		case checkGegex(Module, `^\s*.?help\s*$`):
			fmt.Printf("Comands:\n\t-inventory: loads the inventory module\n\t-grocery: loads the grocery module\n\t-q: exits the program")

		//quit out of program
		case checkGegex(Module, `^\s*.?q\s*$`):
			Checkout()
			os.Exit(1)

		//empty case to prevent default case from running when not using a flag
		case checkGegex(Module, `^\s*.?empty\s*$`):

		//inform of invalid command
		default:
			fmt.Printf("\nCommand, \"%s\" not found. Use, \"help\" for a list of commands\n", Module)
		}

		//Prompt for input
		fmt.Print("Please enter a module to load (inventory or grocery): ")
		fmt.Scanln(&Module)
	}

}