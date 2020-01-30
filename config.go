package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/NehemiahG7/project-0/grocery"
	"github.com/NehemiahG7/project-0/index"
	"github.com/NehemiahG7/project-0/inventory"
)

//Module is the command string for main.go
var Module string

//Inv is pointer to the currently loaded Inventory struct
var Inv *inventory.Inventory

var GrocFile string
var InvFile string
var IdxFile string

//Groc is the currently loaded grocery struct
var Groc *grocery.GroceryList

//Index is the currently loaded ItemIndex
var Index *index.ItemIndex

//CONFIGFILE is a config file
const CONFIGFILE string = "conf.json"

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

	InvFile = config.InvFile
	IdxFile = config.IndexFile
	GrocFile = config.GrocFile

	//load global structs
	Inv = inventory.LoadInv(config.InvFile)
	Groc = grocery.LoadList(*Inv, config.GrocFile)
	Index = index.LoadIndex(*Inv, config.IndexFile)

	//flag for user to enter specifc module
	flag.StringVar(&Module, "module", "menu", "Use this to start the CLI in a specific module. inventory or grocery")
	flag.Parse()

}

// func openFile(str string) *os.File {
// 	file, err := os.Open(str)
// 	if err != nil {
// 		log.Fatalf("File failed to open; %s", err)
// 	}
// 	return file
// }
