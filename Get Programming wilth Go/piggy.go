package main

import (
	"fmt"
)

func main() {
	var piggyBank float64 = 0.0
	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank)
}
