package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

const MAX_SIZE int = 100

func main() {
  fmt.Println(strings.ToUpper("Hello World from Go!")) 

  // Declare and initialize variables  
  var fullName string = "Tahmid Tanzim Lupin"
  fmt.Println(fullName)
  fmt.Printf("variable type = %T\n", fullName)

  var defaultInt int
  fmt.Println(defaultInt)
  fmt.Printf("variable type = %T\n", defaultInt)

  var age int = 26
  fmt.Println(age)
  fmt.Printf("variable type = %T\n", age)

  // Get input from Console
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text - ")
  input, _ := reader.ReadString('\n')
  fmt.Println("You have entered - ", input)

  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
