package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

//Module is the command string for main.go
var Module string

//Inv is pointer to the currently loaded Inventory struct
var Inv *Inventory

//Groc is the currently loaded inventory struct
var Groc *GroceryList

//CONFIGFILE is a config file
const CONFIGFILE string = "conf.json"

//InvFile is a string containing the name of the file to use in invManager
var InvFile string = ""

//GrocFile is a string containing the name of the file to use in groceryList
var GrocFile string = ""

//RecFile is a string containing the name of the file to use in cookbook
var RecFile string = ""

//Configuration is the struct that CONFIGFILE is used to fill
type Configuration struct {
	InvFile  string `json:"InvFile"`
	GrocFile string `json:"GroceryFile"`
	RecFile  string `json:"RecipesFile"`
}

func init() {
	config := Configuration{}
	c, err := os.Open(CONFIGFILE)
	if err != nil {
		log.Fatalf("Failed to open config file: %s", err)
	}
	json.NewDecoder(c).Decode(&config)

	InvFile = config.InvFile
	GrocFile = config.GrocFile
	RecFile = config.RecFile

	Inv = loadInv()
	Groc = createList(*Inv)

	flag.StringVar(&Module, "module", "empty", "Use this to start the CLI in a specific module. inventory, grocery, or cookbook")
	flag.Parse()

}

// func openFile(str string) *os.File {
// 	file, err := os.Open(str)
// 	if err != nil {
// 		log.Fatalf("File failed to open; %s", err)
// 	}
// 	return file
// }
