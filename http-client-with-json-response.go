package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	p := fmt.Println
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
	defer resp.Body.Close()

	p(resp.Status)
	/*scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		p(scanner.Text())
	}*/

	//Unmarshalling JSON
	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &result)
	p(result["menu"])

}
