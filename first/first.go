package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/about", func(resp http.ResponseWriter, req *http.Request) {
		fmt.ServeFile(resp, req, "static/about.html")
	})
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.ServeFile(resp, req, "static/index.html")
	})
	fmt.Println("Server is listening")
	http.ListenAndServe(":8181", nil)
}
