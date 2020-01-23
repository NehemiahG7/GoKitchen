package main

import (
	"fmt"
	"os"
)

func indexModule() {
	for {
		fmt.Print("Welcome to GoKitchen - Index: ")

		commands := Parse(InputString())
		for i := 0; i < len(commands); i++ {
			switch {
			//AddItem
			case checkGegex(commands[i][0], `^\s*.?add\s*$`):
				//Check if user is going to enter items with command
				if checkGegex(commands[i][0], reg) {
					Index.AddItem(commands[i][1], commands[i][2])
					fmt.Println("In Quick Add")
					continue
				}
				//request items
				fmt.Println("What would you like to add?\nPlease seperate all items with a coma:")
				arry := ParseLine(InputString())
				Index.AddItem(arry[0], arry[1])
			//RemoveItem
			case checkGegex(commands[i][0], `^\s*.?remove\s*$`):
				//Check if user is going to enter items with command
				if checkGegex(commands[i][0], reg) {
					Index.RemoveItem(commands[i][1])
					fmt.Println("In Quick Add")
					continue
				}
				//request items
				fmt.Println("What would you like to remove?\nPlease seperate all items with a coma:")
				arry := ParseLine(InputString())
				Index.RemoveItem(arry[0])
			//PrintIndex
			case checkGegex(commands[i][0], `^\s*.?print\s*$`):
				Index.PrintIndex()
			//Help
			case checkGegex(commands[i][0], `^\s*.?help\s*$`):
				fmt.Println(IndexHelpString)
			case checkGegex(commands[i][0], `^\s*.?exit\s*$`):
				Module = "empty"
				Checkout()
				return
			case checkGegex(commands[i][0], `^\s*.?q\s*$`):
				Checkout()
				os.Exit(1)
			default:
				fmt.Printf("Command, \"%s\" not found, continuing with next command", commands[i][0])
			}
		}
	}
}
