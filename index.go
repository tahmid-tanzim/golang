package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	// "math"
)

const MAX_SIZE int = 100

func main() {
  fmt.Println(strings.ToUpper("Hello World from Go!")) 

  // 1. Declare and initialize variables  
  var fullName string = "Tahmid Tanzim Lupin"
  fmt.Println(fullName)
  fmt.Printf("variable type = %T\n", fullName)

  var defaultInt int
  fmt.Println(defaultInt)
  fmt.Printf("variable type = %T\n", defaultInt)

  var age int = 26
  fmt.Println(age)
  fmt.Printf("variable type = %T\n", age)

  // 2. Get input from Console
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text - ")
  input, _ := reader.ReadString('\n')
  fmt.Println("You have entered - ", input)

  // 3. Convert string inputs to other types
  floatString := " 3.86O "
  validConvertedData, err := strconv.ParseFloat(strings.TrimSpace(floatString), 64)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("String to Float - ", validConvertedData)
  }

  // 4. Convert types before using
  var a int = 8
  var b float64 = 42.75
  ab0 := float64(a) + b
  ab1 := a + int(b)
  fmt.Printf("(float64) Total - %v, Type - %T\n", ab0, ab0)
  fmt.Printf("(int) Total - %v, Type - %T\n", ab1, ab1)

  // 5. Use the math package
  c1, c2, c3 := 23.5, 65.1, 76.3
  c0 := c1 + c2 + c3
  fmt.Println(c0)

  c0 = math.Round(c0 * 100) / 100
  fmt.Println(c0)

  circleRadius := 15.5
  circumference := circleRadius * 2 * math.Pi
  fmt.Printf("Circumference - %.2f\n", circumference)

  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
