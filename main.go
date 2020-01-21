package main

import (
	"fmt"
	"os"
)

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
		switch Module {
		case "print":
			Inv.printInv()
		case "help":
			fmt.Printf("Comands:\nprint: Prints current inventory\n")
		case "exit":
			Module = "empty"
			return
		case "q":
			os.Exit(1)
		default:
			fmt.Print("Welcome to GoKitchen - Inventory: ")
		}
		fmt.Scanln(&Module)
	}
}
func groceryModule() {
	for {
		switch Module {
		case "print":
			Groc.Print()
		case "help":
			fmt.Print("Comands:\nprint: Prints current grocery list\n")
		case "exit":
			Module = "empty"
			return
		case "q":
			os.Exit(1)
		default:
			fmt.Print("Welcome to GoKitchen - Grocery: ")
		}
		fmt.Scanln(&Module)
	}
}
