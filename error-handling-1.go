package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	unIdenfifiedObj error = errors.New("Unable to identify object..")
)

func main() {
	err := egFunc()
	if err != nil {
		fmt.Println(unIdenfifiedObj)
	}

}

func egFunc() error {
	_, err := os.Open("/non-existent")
	if err != nil {
		return unIdenfifiedObj
	}
	return nil
}
