package main

import (
	"strings"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/NehemiahG7/GoKitchen/inventory"
	"github.com/NehemiahG7/GoKitchen/util"
)

var nm string = "Everyone"
var reg string = `^.*\?.*$`

type account struct{
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
	acc := account{
		Name: "stuff",
	}
	if r.Method == "POST"{
		fmt.Printf("%s\n", r.PostFormValue("account"))
		http.Redirect(w, r, "http://localhost:8080/inv", http.StatusSeeOther)
	}
	servHTML(w, "html/login.html", acc)

}

func inv(w http.ResponseWriter, r *http.Request){

	acc := account{
		Name: nm,
		Inv: *inventory.LoadInv(InvFile),
	}

	//if get request contains /?, process form request
	if util.CheckGegex(r.RequestURI, reg){
		formResp(r.RequestURI, acc)
		fmt.Printf("Processing form response\n")
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

func formResp(str string, acc account){

	//Trim prefix from str and split into an arry containing the input
	str = strings.TrimPrefix(str, "/inv?")
	arry := strings.Split(str, "=")

	//Determine which request was made and exicute
	if arry[0] == "item"{

		arry[1] = strings.TrimSuffix(arry[1], "&cat")
		arry = arry[1:3]
		temp:=arry[1]
		arry[1] = arry[0]
		arry[0] = temp

		if arry[1] == ""{
			return
		}
		acc.Inv.Add(arry)

		fmt.Printf("1st entry %s, 2nd entry %s\n", arry[0], arry[1])
	} else if arry[0] == "addCat"{
		arry = arry[1:]
		if arry[0] == ""{
			return
		}
		acc.Inv.Add(arry)
	}

	util.Encode(acc.Inv, InvFile)
}

func startTCP(content string, address string, cha chan string){
	// con, err := net.Dial("TCP", address)
	// if err != nil {
	// 	fmt.Printf("Dial err: %s\n", err)
	// }
	//con.Write(content)
}