package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	p := fmt.Println
	file := "resources/sample.log"
	data, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	i := 0
	ips := 0

	errorIPs := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			i++
			words := strings.Split(line, " ")
			for _, word := range words {
				isIP := net.ParseIP(word)
				if isIP.To4() != nil || isIP.To16() != nil {
					errorIPs[word] = true
					ips++
				}
			}
		}
	}
	p("ERROR Logs, IPs: ", i, ips)
	p("IPs :", errorIPs)
}

/*
Output:

ERROR Logs, IPs:  5 4
IPs : map[10.46.12.56:true 192.168.0.0:true 192.46.12.56:true FE80:0000:0000:0000:0202:B3FF:FE1E:8329:true]
*/
