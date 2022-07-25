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
