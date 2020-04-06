package main

import (
	"net/url"
	"fmt"
	"html/template"
	"log"
	"strconv"
	"net/http"
	"github.com/NehemiahG7/GoKitchen/internal/inventory"
	"github.com/NehemiahG7/GoKitchen/internal/account"
)

var reg string = `^.*\?.*$`

type accStruc struct{
	ID int
	Name string
	Inv inventory.Inventory
}

func main(){

	http.HandleFunc("/", login)
	http.HandleFunc("/inv", inv)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("html/static/"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request){

	_ , err := r.Cookie("id")
	if err == nil{
		http.Redirect(w, r, "http://localhost:8080/inv", http.StatusSeeOther)
		return
	}

	if r.Method == "POST"{
		id := account.Login(r.PostFormValue("account"), r.PostFormValue("password"))
		if id != 0 {
			cookie := http.Cookie{
				Name: "id",
				Value: strconv.Itoa(id),
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "http://localhost:8080/inv", http.StatusSeeOther)
			return
		} 
		if id == 0 {
			id = account.CreateAccount(r.PostFormValue("account"), r.PostFormValue("password"))
			cookie := http.Cookie{
				Name: "id",
				Value: strconv.Itoa(id),
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "http://localhost:8080/inv", http.StatusSeeOther)
			return
		}
	}
	acc := accStruc{}
	servHTML(w, "html/login.html", acc)

}

func inv(w http.ResponseWriter, r *http.Request){

	//If user doesn't have the id cookie, redirect them to login
	//Otherwise, get cooke and convert to useable type
	cookie, err := r.Cookie("id")
	if err != nil{
		fmt.Printf("invHandler: redirecting user to login, %s\n", err)
		http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
		return
	}
	id, err := strconv.Atoi(cookie.Value)
	if err != nil{
		fmt.Printf("invHandler: error parsing value, %s\n", err)
	}

	//Get account info
	acc := accStruc{
		ID: id,
		Name: account.GetUsername(id),
		Inv: account.GetInv(id),
		//Inv: *inventory.LoadInv(InvFile),
	}

	query := r.URL.RawQuery
	if query != ""{
		formResp(query, acc)
	}

	servHTML(w, "html/index.html", acc)
}

func servHTML(w http.ResponseWriter, file string, stuc interface{}){
	//Parse html template to serve
	landing, err := template.ParseFiles(file)
	if err != nil{
		log.Print(err)
	}
	
	//serve html template and acc struct
	err = landing.Execute(w, stuc)
	if err != nil{
		log.Print(err)
	}
}

func formResp(query string, acc accStruc){
	values, err := url.ParseQuery(query)
	if err != nil{
		fmt.Printf("formResp: error parsing query: %s\n", err)
	}
	cat, ok := values["addCat"]
	if ok {
		account.AddCatagory(cat[0], acc.Inv, acc.ID)
		return
	}
	item, ok := values["item"]
	if ok {
		cat, _ := values["cat"]
		account.AddItem(item[0],cat[0], acc.Inv, acc.ID)
	}
}