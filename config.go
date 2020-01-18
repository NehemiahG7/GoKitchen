package main

import (
	"encoding/json"
	"log"
	"os"
)

const CONFIGFILE string = "conf.json"

var InvFile string = ""
var GrocFile string = ""
var RecFile string = ""

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

}

// func openFile(str string) *os.File {
// 	file, err := os.Open(str)
// 	if err != nil {
// 		log.Fatalf("File failed to open; %s", err)
// 	}
// 	return file
// }
