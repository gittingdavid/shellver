package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9000", nil) //Set port
}

func login(response http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method) //Get request method
	if request.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(response, nil)
	} else {
		request.ParseForm()
		// Print form input to terminal
		//fmt.Println("Username:", r.Form["username"])
		//fmt.Println("Password:", r.Form["password"])
		//fmt.Println("IP Address:", r.Form["ip"])

		var username string = fmt.Sprint(request.Form["username"])
		var password string = fmt.Sprint(request.Form["password"])
		var ip string = fmt.Sprint(request.Form["ip"])

		connect(username, password, ip)

		output := []byte("Loading. . . ")
		fmt.Println("Page is loading")
		response.Write(output)
	}
}

func connect(username string, password string, ip string) {
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
	fmt.Println("IP Address:", ip)

	// Connect to ssh cient
	config := &ssh.ClientConfig{
		//To resolve "Failed to dial: ssh: must specify HostKeyCallback" error
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		//panic("Dial Failed: " + err.Error())
		fmt.Println("Invalid Login or Password")
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Session Failed: " + err.Error())
	}
	defer session.Close()

	/////////////////////////////////////////////////////////////

	fmt.Println("Successfully Connected!")

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Shell()
	session.Wait()

}
