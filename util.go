package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func encode(item interface{}, str string) {

	file, err := os.Create(str)
	if err != nil {
		log.Fatalf("File failed to create %s", err)
	}
	defer file.Close()

	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.Encode(item)

	_, err = io.Copy(file, buf)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

//Parse takes a string of user input, and returns [][]string where the first entry of each [] is a command for the app
func Parse(str string) [][]string {
	strArry := [][]string{}
	strs := strings.Split(str, "-")

	for i := 0; i < len(strs); i++ {
		strArry = append(strArry, strings.Split(strs[i], ","))
	}
	return strArry
}

//ParseLine takes a string of user input and returns []string
func ParseLine(str string) []string {
	return strings.Split(str, ",")
}
func checkGegex(str, rg string) bool {
	b, err := regexp.MatchString(rg, str)
	if err != nil {
		return false
	}
	return b
}

//Checkout encodes all currently open structs to their respective files
func Checkout() {
	Groc.encodeList()
	Inv.encodeInv()
}
