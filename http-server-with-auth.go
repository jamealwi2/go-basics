package main

import (
	"fmt"
	"net/http"
)

/*
curl -v -X POST http://localhost:8081/hello -H "Authorization: Basic YWJjOjEyMw=="
*/

func hello(w http.ResponseWriter, req *http.Request) {
	u, p, ok := req.BasicAuth()
	if !ok {
		w.WriteHeader(401)
		return
	}

	if u != "abc" {
		fmt.Fprintf(w, "Incorrect uname")
		w.WriteHeader(401)
		return
	}

	if p != "123" {
		fmt.Fprintf(w, "Incorrect pass")
		w.WriteHeader(401)
		return
	}

	fmt.Fprintf(w, "Successfully Logged In")
	w.WriteHeader(200)

}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8081", nil)

}
