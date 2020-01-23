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

//Groc is the currently loaded grocery struct
var Groc *GroceryList

//Index is the currently loaded ItemIndex
var Index *ItemIndex

//CONFIGFILE is a config file
const CONFIGFILE string = "conf.json"

//InvFile is a string containing the name of the file to use in invManager
var InvFile string = ""

//GrocFile is a string containing the name of the file to use in groceryList
var GrocFile string = ""

//IndexFile is a string containint the name of the file used in itemIndex
var IndexFile string = ""

//RecFile is a string containing the name of the file to use in cookbook
//var RecFile string = ""

//Configuration is the struct that CONFIGFILE is used to fill
type Configuration struct {
	InvFile   string `json:"InvFile"`
	GrocFile  string `json:"GroceryFile"`
	IndexFile string `json:"IndexFile"`
	//RecFile   string `json:"RecipesFile"`

}

func init() {
	//get config file
	config := Configuration{}
	c, err := os.Open(CONFIGFILE)
	if err != nil {
		log.Fatalf("Failed to open config file: %s", err)
	}
	json.NewDecoder(c).Decode(&config)

	//assign file names to global veriables
	InvFile = config.InvFile
	GrocFile = config.GrocFile
	IndexFile = config.IndexFile
	//RecFile = config.RecFile

	//load global structs
	Inv = loadInv()
	Groc = loadList()
	Index = LoadIndex()

	//flag for user to enter specifc module
	flag.StringVar(&Module, "module", "empty", "Use this to start the CLI in a specific module. inventory or grocery")
	flag.Parse()

}

// func openFile(str string) *os.File {
// 	file, err := os.Open(str)
// 	if err != nil {
// 		log.Fatalf("File failed to open; %s", err)
// 	}
// 	return file
// }
