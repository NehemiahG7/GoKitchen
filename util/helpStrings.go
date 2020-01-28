package util

//InventoryHelpString is the string used in main.go for the help command in Inventory module
var InventoryHelpString string = `Help
-Putting a "/" in front of any input command allows for input to be issued with the command
-If a "/" is not provided, there is will be a prompt for input
-Multiple commands can be entered at once by putting a "-" before the command
ex: -/add,meats,sausage,beef,chicken -/remove,sausage,beef,chicken
Commands:
	-print: Prints current inventory
	-add: Adds the items to the inventory
		*Items can only be added to one catagory at a time
		*Seperate each catagory and item by a coma
		*If a new catagory is used, it will make a new group in the list
		ex: /add, meats, chicken, beef
	-addCatagory: Adds the given catagory to the inventory
		*Command can only add one catagory per call
		*Seperate command and catagory with a coma
		ex: /addCatagory,frozen
	-removeCatagory: Removes the given catagory from the inventory
		*Command can only rmove one catagory per call
		*Caution* this will remove all items in the given catagory
		*Sperate command and catagory with a coma
	-remove: Removes the provided items from the inventory
		*No catagory is needed
		*Seperate each item with a coma
		ex: /remove, chicken, beef
	-addToGrocery: Adds the provided items to the grocery list
		*No catagory is needed
		*This will only add items that are already in the inventory
		ex: /addToGrocery, chicken, beef
	-removeFromGrocery: Removes the provided items from the grocery list
		*No catagory is needed
		*Only removes items that are already in the inventory
		ex: /removeFromGrocery, chicken, beef
	-changeCatagory: Moves an item from one catagory to another
		*Items must be entered in pairs with the new catagory coming before the item
		ex: /changeCatagory, meats, chicken, fruits, apples, fruits, bannannas
	-exit: Returns to the main menu
	-q: Closes the program`

//GroceryHelpString is the string used in main.go for the help command in Grocery module
var GroceryHelpString string = `Help
-Putting a "/" in front of any input command allows for input to be issued with the command
-If a "/" is not provided, there is will be a prompt for input
-Multiple commands can be entered at once by putting a "-" before the command
ex: -/add, chicken, beef, onions-/remove, chicken, beef, onions-print
Commands:
	-print: Prints current grocery list
	-add: Adds the items to the grocery list
		*Seperate each catagory and item by a coma
		ex: /add, chicken, beef, onions
	-remove: Removes the provided items from the grocery list
		*No catagory is needed
		*Seperate each item with a coma
		ex: /remove, chicken, beef
	-addToInv: Adds all items in the grocery list to the inventory
		*Catagories are retrieved from an index of all items entered for this inventory
		*If an item has not been entered, it will be assigned to the "other" catagory
		*If an item already exists, its date will be updated
		ex: addToInv
	-export: saves the grocery list to groceryList.txt
		ex: export
	-index: Loads the Index sub-module
	-exit: returns to the main menu
	-q: closes the program`

//IndexHelpString is the string used in main.go for the help command in Index module
var IndexHelpString string = `Help
-Putting a "/" in front of any input command allows for input to be issued with the command
-If a "/" is not provided, there is will be a prompt for input
-Multiple commands can be entered at once by putting a "-" before the command
ex: -/add, chicken, beef, onions-/remove, chicken, beef, onions-print
Commands:
	-print: Prints current index
	-add: Adds the items to the index
		*Input must have the name followed by the catagory
		*Seperate each catagory and item by a coma
		*This command does NOT allow for multiple inputs at once
		ex: /add, chicken, meats
	-remove: Removes the provided items from the index
		*Just the name is needed
		*This command does NOT allow for multiple inputs at once
		ex: /remove, chicken
	-exit: returns to the main menu
	-q: closes the program`
