package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var Port = ":8080"

/**
*Author:KITAGAWA YUTO
*DATE:4/12
**/

//main function
func main() {
	//if pattern is "/", ServerFile function will run
	http.HandleFunc("/", ServeFiles)
	//it is log in the terminal
	fmt.Println("Serving @ : ", "http://127.0.0.1"+Port)
	//nil means confortable
	log.Fatal(http.ListenAndServe(Port, nil))
}

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//For example, search something and url added by get method
	case "GET":
		path := r.URL.Path
		fmt.Println(path)
		if path == "/" {
			path = "./static/index.html"
		} else {
			path = "." + path
		}

		http.ServeFile(w, r, path)

		//For example, submited something by input
	case "POST":
		r.ParseMultipartForm(0)
		fmt.Println("come someone")
		//it is id="message"
		message := r.FormValue("message")
		//fmt run in the terminal
		fmt.Println("----------------------")
		fmt.Println("Message from Client:", message)
		fmt.Fprintf(w, "Server: %s \n ", message+" | "+time.Now().Format(time.RFC3339))

		//other method neither POST nor GET
	default:
		fmt.Println(w, "Request type other than GET or POST not supported")
	}

}
