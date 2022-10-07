package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/hello", nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("abc", "123")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
