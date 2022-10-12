package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	p := fmt.Println
	filePath := "resources/sample.log"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	//Reference time -> Mon Jan 2 15:04:05 MST 2006
	startTime, err := time.Parse("15:04:05.999999", "02:22:14.110")
	endTime, err := time.Parse("15:04:05.999999", "02:22:19.300")

	errors := 0

	for scanner.Scan() {
		log := scanner.Text()
		if strings.Contains(log, "ERROR") {
			words := strings.Split(log, " ")
			for _, word := range words {
				t, err := time.Parse("15:04:05.999999", word)
				if err == nil {
					if t.After(startTime) && t.Before(endTime) {
						errors++
						p(word)
					}
				}
			}
		}
	}
	p(errors)
}

/*
Output:
02:22:19.190
02:22:19.218
02:22:19.218
3
*/
