package inventory

import (
	"fmt"
	"os"

	_ "github.com/NehemiahG7/project-0/util"
)

func inventoryModule(Inv Inventory) {
	for {
		fmt.Print("Welcome to GoKitchen - Inventory: ")
		commands := Parse(InputString())

		for i := 0; i < len(commands); i++ {
			switch {
			//Add single or multiple items to the Kitchen Inventory
			case CheckGegex(commands[i][0], `^\s*.?add\s*$`):
				//Make an arry to hold input []string
				var arry = make([]string, 0)
				//Check if user is going to enter items with command
				if CheckGegex(commands[i][0], reg) {
					arry = commands[i][1:]
				} else {
					//request items
					fmt.Println("What would you like to add? \nPlease format enter a catagory followed by a coma, \nand then any item you wish seperated by only comas:")
					arry = ParseLine(InputString())
				}
				//Check if the given key is currently a key in Inv
				_, b := Inv.Inven[arry[0]]
				//If it is not, prompt user if they'd like to make it
				if !b {
					fmt.Printf("Could not find catagory, '%s' would you like to add it(y/n)? ", arry[0])
					str := InputString()
				Loop:
					for {
						switch str {
						//if y, add the input
						case "y":
							Inv.Add(arry)
							break Loop
						//if n, get the correct key
						case "n":
							fmt.Printf("What catagory would you like to add it to? ")
							str2 := InputString()
							arry[0] = str2
							Inv.Add(arry)
							break Loop
						//make sure to get valid input
						default:
							fmt.Printf("Please repeat: (y/n)")
							str = InputString()
						}
					}
					//add input if valid key was present
				} else {
					Inv.Add(arry)
				}

			//Adds a catagory to the Inventory
			case checkGegex(commands[i][0], `^\s*.?addCatagory\s*$`):
				if checkGegex(commands[i][0], reg) {
					Inv.Add(commands[i][1:])
				}
				fmt.Print("What catagory would you like to add: ")
				Inv.Add(ParseLine(InputString()))

			//Removes a catagory from the Inventory, including it's []item
			case checkGegex(commands[i][0], `^\s*.?removeCatagory\s*$`):
				if checkGegex(commands[i][0], reg) {
					Inv.RemoveKey(commands[i][1])
				}
				fmt.Print("What catagory would you like to remove? (caution, this will remove all items in that catagory): ")
				arry := ParseLine(InputString())
				Inv.RemoveKey(arry[0])
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

				Inv.Remove(ParseLine(InputString()))

			//Toggle Inv.ForceGrocery to true for specified items
			case checkGegex(commands[i][0], `^\s*.?addToGrocery\s*$`):
				//Check if user is entering items with command
				if checkGegex(commands[i][0], reg) {
					Inv.AddGrocery(commands[i][1:])
					continue
				}
				//Prompt user for input
				fmt.Println("What would you like to add to the grocery list(Note: this only adds items that are already in Inventory)")

				Inv.AddGrocery(ParseLine(InputString()))

			//Toggle Inv.ForceGrocery to false
			case checkGegex(commands[i][0], `^\s*.?removeFromGrocery\s*$`):
				//Check if user is entering items with command
				if checkGegex(commands[i][0], reg) {
					Inv.RemoveGrocery(commands[i][1:])
					continue
				}
				//Prompt user for input
				fmt.Println("What would you like to remove from the grocery list?(Note: this only removes items that are already in Inventory)")

				Inv.RemoveGrocery(ParseLine(InputString()))

			//Move the given item to a new key
			case checkGegex(commands[i][0], `^\s*.?changeCatagory\s*$`):
				if checkGegex(commands[i][0], reg) {
					Inv.ChangeKey(commands[i][1:])
					continue
				}
				fmt.Println("What would you like to edit?(Note: Enter all pairs key first, and seperated only by comas")

				Inv.ChangeKey(ParseLine(InputString()))

			//Print the current kitchen inventory
			case checkGegex(commands[i][0], `^\s*.?print\s*$`):
				Inv.printInv()

			//Print list of commands
			case checkGegex(commands[i][0], `^\s*.?help\s*$`):
				fmt.Println(InventoryHelpString)

			//Exit to main screen
			case checkGegex(commands[i][0], `^\s*.?exit\s*$`):
				Module = "empty"
				Checkout()
				return

			//Quit app
			case checkGegex(commands[i][0], `^\s*.?q\s*$`):
				Checkout()
				os.Exit(1)

			//Invalid command
			default:
				fmt.Printf("Command \"%s\" not found, continuing to next command\n", Module)
			}
		}

	}
}
