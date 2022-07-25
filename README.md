# RBL Checker

The RBL package checks the IP Address in the Real-time Blackhole List - a list where malicious addresses of email clients are entered.

Example:
```go
package main

import (
	"github.com/qwddz/rbl"
	"log"
)

func main() {
	res, err := rbl.New().CheckIPAddress("91.240.141.42")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Status, res.Errors)
}
```
