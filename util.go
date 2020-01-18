package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func encode(item interface{}, str string) {

	file, err := os.Create(str)
	if err != nil {
		log.Fatalf("File failed to create %s", err)
	}

	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.Encode(item)

	_, err = io.Copy(file, buf)
	if err != nil {
		fmt.Printf("%s", err)
	}

	file.Close()
}
