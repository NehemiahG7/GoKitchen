package main

import (
	"log"
	"net/http"
	
)

func main(){

	http.HandleFunc("/", login)
	http.HandleFunc("/inv", inv)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}

}

func login(w http.ResponseWriter, r *http.Request){

}

func inv(w http.ResponseWriter, r *http.Request){

}