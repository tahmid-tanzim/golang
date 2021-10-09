package main

import (
	"fmt"
	"oneleven.com/greetings"
)

func main() {
	message := greetings.Hello("Tanzim")
	fmt.Println(message)
}