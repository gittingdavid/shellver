package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hellohtml", hellohtml)
	http.ListenAndServe(":9000", nil)
	// Use browser and go to "localhost:9000

}

func hello(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path[1:])
	output := []byte("Hello There!")
	fmt.Println("Someone says hello!")
	response.Write(output)
	//Use browser and go to "localhost:9000/hello"
}

func hellohtml(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")

	/*
		output := []byte("<html><body><h1>Hello There!</h1></body></html>")
		response.Write(output)
	*/

	io.WriteString(response, `
	<DOCTYPE html>
	<html>
	<head>
		<title> My Page </title>
	</head>
	<body>
		<h2> Welcome to my page </h2>
		<p> This is a test of a go server <p>
	</body>
	</html>
	`)
}
