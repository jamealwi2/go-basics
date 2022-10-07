package main

import (
	"fmt"
	"net/http"
)

var (
	jsonResp string = `{"menu": {
		"id": "file",s
		"value": "File",
		"popup": {
		  "menuitem": [
			{"value": "New", "onclick": "CreateNewDoc()"},
			{"value": "Open", "onclick": "OpenDoc()"},
			{"value": "Close", "onclick": "CloseDoc()"}
		  ]
		}
	  }}`
)

func hello(w http.ResponseWriter, req *http.Request) {
	u, p, ok := req.BasicAuth()
	if !ok {
		w.WriteHeader(401)
	}
	if u != "abc" {
		w.WriteHeader(401)
	}
	if p != "123" {
		w.WriteHeader(401)
	}
	fmt.Fprintf(w, jsonResp)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8081", nil)

}
