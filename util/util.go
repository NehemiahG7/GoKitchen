package util

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

//Encode exports given items to the given file name,str
func Encode(item interface{}, str string) {

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

//InputString gets user from the input and returns it all as a string
func InputString() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()

}

//Parse takes a string of user input, and returns [][]string where the first entry of each [] is a command for the app
func Parse(str string) [][]string {
	rgx := regexp.MustCompile(`\s*-\s*`)
	strs := rgx.Split(str, -1)
	strArry := [][]string{}

	for i := 0; i < len(strs); i++ {
		strArry = append(strArry, ParseLine(strs[i]))
	}
	return strArry
}

//ParseLine takes a string of user input and returns []string
func ParseLine(str string) []string {
	rgx := regexp.MustCompile(`\s*,\s*`)

	return rgx.Split(str, -1)
}

//CheckGegex takes a string and a regex expression in a the form of a string and returns true if they match, false if they don't
func CheckGegex(str, rg string) bool {
	b, err := regexp.MatchString(rg, str)
	if err != nil {
		return false
	}
	return b
}
