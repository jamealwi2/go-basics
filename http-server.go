package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, res *http.Request) {
	fmt.Println("hello")
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8081", nil)
}
