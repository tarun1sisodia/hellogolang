package main

import (
	"fmt"
	"log"

	"github.com/tarun1siodia/Golang"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Tarun")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
