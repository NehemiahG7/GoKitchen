package main

import (
	"fmt"
	"os"
)

var reg string = `/.*`

func main() {

	for {
		switch Module {
		case "inventory":
			inventoryModule()
		case "grocery":
			groceryModule()
		case "help":
			fmt.Printf("Comands:\nprint: Prints current inventory\n")
		case "q":
			os.Exit(1)
		default:
			fmt.Print("Please enter a module to load (inventory, grocery, cookbook): ")
		}
		fmt.Scanln(&Module)

	}

}
func inventoryModule() {
	for {
		fmt.Print("Welcome to GoKitchen - Inventory: ")
		fmt.Scanln(&Module)
		commands := Parse(Module)

		for i := 0; i < len(commands); i++ {
			Module = commands[i][0]
			switch {
			//Add single or multiple items to the Kitchen Inventory
			case checkGegex(commands[i][0], `^\s*.?add\s*$`):
				//Check if user is going to enter items with command
				if checkGegex(commands[i][0], reg) {
					Inv.Add(commands[i][1:])
					fmt.Println("In Quick Add")
					continue
				}
				//request items
				fmt.Println("What would you like to add? \nPlease format enter a catagory(meats,fruits,vegetables,grains,dairy) followed by a coma, \nand then any item you wish seperated by only comas:")
				fmt.Scanln(&Module)
				Inv.Add(ParseLine(Module))

			//Reinitilize the entire inventory
			case checkGegex(commands[i][0], `^\s*.?reInIt\s*$`):
				Inv = createInv()

			//Remove single or multiple items from the inventory
			case checkGegex(commands[i][0], `^\s*.?remove\s*$`):
				//Check if use is entering items with the command
				if checkGegex(commands[i][0], reg) {
					Inv.Remove(commands[i][1:])
					continue
				}
				//Prompt use for items to remove
				fmt.Println("What would you like to remove?\nPlease enter all items sperated by only comas: ")
				fmt.Scanln(&Module)
				Inv.Remove(ParseLine(Module))

			//Toggle Inv.ForceGrocery to true for specified items
			case checkGegex(commands[i][0], `^\s*.?addToGrocery\s*$`):
				//Check if user is entering items with command
				if checkGegex(commands[i][0], reg) {
					Inv.AddGrocery(commands[i][1:])
					continue
				}
				//Prompt user for input
				fmt.Println("What would you like to add to the grocery list(Note: this only adds items that are already in Inventory)")
				fmt.Scanln(&Module)
				Inv.AddGrocery(ParseLine(Module))

			//Toggle Inv.ForceGrocery to false
			case checkGegex(commands[i][0], `^\s*.?removeFromGrocery\s*$`):
				//Check if user is entering items with command
				if checkGegex(commands[i][0], reg) {
					Inv.RemoveGrocery(commands[i][1:])
					continue
				}
				//Prompt user for input
				fmt.Println("What would you like to remove fro mthe grocery list?(Note: this only removes items that are already in Inventory)")
				fmt.Scanln(&Module)
				Inv.RemoveGrocery(ParseLine(Module))

			//Move the given item to a new key
			case checkGegex(commands[i][0], `^\s*.?changeKey\s*$`):
				if checkGegex(commands[i][0], reg) {
					Inv.ChangeKey(commands[i][1:])
					continue
				}
				fmt.Println("What would you like to edit?(Note: Enter all pairs key first, and seperated only by comas")
				fmt.Scanln(&Module)
				Inv.ChangeKey(ParseLine(Module))

			//Print the current kitchen inventory
			case checkGegex(commands[i][0], `^\s*.?print\s*$`):
				Inv.printInv()

			//Print list of commands
			case checkGegex(commands[i][0], `^\s*.?help\s*$`):
				fmt.Printf("Comands:\nprint: Prints current inventory\n")

			//Exit to main screen
			case checkGegex(commands[i][0], `^\s*.?exit\s*$`):
				Module = "empty"
				Inv.encodeInv()
				return

			//Quit app
			case checkGegex(commands[i][0], `^\s*.?q\s*$`):
				Inv.encodeInv()
				os.Exit(1)

			//Invalid command
			default:
				fmt.Printf("Command \"%s\" not found, continuing to next command\n", Module)
			}
		}

	}
}
func groceryModule() {
	for {
		fmt.Print("Welcome to GoKitchen - Grocery: ")
		fmt.Scanln(&Module)
		commands := Parse(Module)
		for i := 0; i < len(commands); i++ {
			switch {
			case checkGegex(commands[i][0], `.*print`):
				Groc.Print()
			case checkGegex(commands[i][0], `.*help`):
				fmt.Print("Comands:\nprint: Prints current grocery list\n")
			case checkGegex(commands[i][0], `.*exit`):
				Module = "empty"
				return
			case checkGegex(commands[i][0], `q`):
				os.Exit(1)
			default:
				fmt.Printf("Command, \"%s\" not found, continuing with next command", commands[i][0])

			}
		}
	}
}
