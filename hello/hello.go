package main

import (
	"fmt"
	"log"
	"oneleven.com/greetings"
)

func main() {
	log.SetPrefix("greetings - ")
	log.SetFlags(0)

	message, err := greetings.Hello("Tahmid Tanzim")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}