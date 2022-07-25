package main

import (
	"github.com/qwddz/rbl"
	"log"
)

func main() {
	log.Println("start checking...")

	res, err := rbl.New().CheckIP("3.124.8.148")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Status, res.Errors)
}
