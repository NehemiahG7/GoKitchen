package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/NehemiahG7/GoKitchen/inventory"
)

type account struct{
	name string
	inv inventory.Inventory
}

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