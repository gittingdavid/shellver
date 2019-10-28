package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9000", nil) //Set port
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //Get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// Print form input to terminal
		fmt.Println("Username:", r.Form["username"])
		fmt.Println("Password:", r.Form["password"])
		fmt.Println("IP Address:", r.Form["ip"])
	}
}
