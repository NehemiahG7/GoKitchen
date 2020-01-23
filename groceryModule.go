package main

import (
	"fmt"
	"os"
)

func groceryModule() {
	for {
		fmt.Print("Welcome to GoKitchen - Grocery: ")

		commands := Parse(InputString())
		for i := 0; i < len(commands); i++ {
			switch {
			//Add
			case checkGegex(commands[i][0], `^\s*.?add\s*$`):
				//Check if user is going to enter items with command
				if checkGegex(commands[i][0], reg) {
					Groc.Add(commands[i][1:])
					continue
				}
				//request items
				fmt.Println("What would you like to add?\nPlease seperate all items with a coma:")
				Groc.Add(ParseLine(InputString()))

			//Remove
			case checkGegex(commands[i][0], `^\s*.?remove\s*$`):
				//Check if user is going to enter items with command
				if checkGegex(commands[i][0], reg) {
					Groc.Remove(commands[i][1:])
					continue
				}
				//request items
				fmt.Println("What would you like to remove?\nPlease seperate all items with a coma:")
				Groc.Remove(ParseLine(InputString()))

			//AddToInv
			case checkGegex(commands[i][0], `^\s*.?addToInv\s*$`):
				Groc.AddToInv()
				fmt.Println("Inventory updated from grocery list")

			//ExportGrocery
			case checkGegex(commands[i][0], `^\s*.?export\s*$`):
				Groc.ExportList()
				fmt.Println("Grocery list saved to groceryList.txt")

			//Print the grocery list
			case checkGegex(commands[i][0], `^\s*.?print\s*$`):
				Groc.Print()

			//help
			case checkGegex(commands[i][0], `^\s*.?help\s*$`):
				fmt.Println(GroceryHelpString)

			//move to index sub-module
			case checkGegex(commands[i][0], `^\s*.?index\s*$`):
				Module = "index"
				Checkout()
				indexModule()
				return

			//return to main
			case checkGegex(commands[i][0], `^\s*.?exit\s*$`):
				Module = "empty"
				Checkout()
				return

			//quit program
			case checkGegex(commands[i][0], `^\s*.?q\s*$`):
				Checkout()
				os.Exit(1)

			//invalid command
			default:
				fmt.Printf("Command, \"%s\" not found, continuing with next command", commands[i][0])
			}
		}
	}
}
