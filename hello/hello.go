package main

import (
	"fmt"
	"log"
	"oneleven.com/greetings"
)

func main() {
	log.SetPrefix("greetings - ")
	log.SetFlags(0)

	names := []string{
		"Lupin",
		"Fatiha",
		"Serhan",
	}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}