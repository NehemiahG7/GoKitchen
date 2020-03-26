package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/NehemiahG7/GoKitchen/inventory"
)

var nm string = "salty"

type account struct{
	Name string
	Inv inventory.Inventory
}

func main(){

	http.HandleFunc("/log", login)
	http.HandleFunc("/", inv)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("html/static/"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request){

}

func inv(w http.ResponseWriter, r *http.Request){

	acc := account{
		Name: nm,
		Inv: *inventory.LoadInv(InvFile),
	}

	landing, err := template.ParseFiles("html/index.html")
	if err != nil{
		log.Print(err)
	}
	
	err = landing.Execute(w, acc)
	if err != nil{
		log.Print(err)
	}
}