package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const MAX_SIZE int = 100

type Dog struct {
	Breed  string
	Weight int
}

func (d Dog) toString() {
	fmt.Println("My Breed is", d.Breed)
}

const URL = "https://gorest.co.in/public/v1/users"

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
		// panic(err)
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

	c0 = math.Round(c0*100) / 100
	fmt.Println(c0)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference - %.2f\n", circumference)

	// 6. Work with dates and times
	n := time.Now()
	fmt.Println("Date & Time - ", n.Format(time.ANSIC))

	// 7. How memory is allocated & managed
	student := make(map[string]int)
	student["history"] = 35
	student["biology"] = 75
	fmt.Println("Student -", student)

	// 8. Reference values with pointers
	var d int = 42
	var pointer *int = &d
	fmt.Println("Before value & pointer - ", d, *pointer)
	*pointer = 31
	fmt.Println("After value & pointer - ", d, *pointer)

	// 9. Store ordered values in arrays
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)

	var numbers = [3]int{105, 425}
	fmt.Println(numbers, len(numbers))

	// 10. Manage ordered values in slices
	var colours = []string{"Red", "Green", "Blue"}
	colours = append(colours, "White")
	fmt.Println(colours)

	// Remove 1st & last item from slice
	colours = append(colours[1 : len(colours)-1])
	fmt.Println(colours)

	numbersSlice := make([]int, 5)
	numbersSlice[0] = 32
	numbersSlice[1] = 43
	numbersSlice[2] = 76
	numbersSlice[3] = 21
	numbersSlice[4] = 87
	numbersSlice = append(numbersSlice, 13)
	fmt.Println("Before sort -", numbersSlice)
	sort.Ints(numbersSlice)
	fmt.Println("After sort -", numbersSlice)

	// 11. Store unordered values in maps
	states := make(map[string]string)
	fmt.Println(states)
	states["BC"] = "British Columbia"
	states["ON"] = "Ontario"
	states["SK"] = "Saskatchwan"
	for key, val := range states {
		fmt.Printf("%v: %v\n", key, val)
	}
	delete(states, "SK")
	fmt.Println(states)

	// 12. Group related values in structs
	poodle := Dog{"Poodle", 12}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	poodle.Weight += 1
	fmt.Printf("Breed:%v Weight:%v\n", poodle.Breed, poodle.Weight)
	poodle.toString()

	// 13. Program conditional logic
	var e int = 3
	var result string
	if e > 0 {
		result = "Greater than Zero"
	} else if e < 0 {
		result = "Less than Zero"
	} else {
		result = "Equal to Zero"
	}
	fmt.Println(result)

	// 14. Evaluate expressions with switch statements
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)
	switch dow {
	case 1:
		result = "Sunday"
	case 2:
		result = "Monday"
	case 3:
		fallthrough
	case 4:
		result = "Tuesday & Wednesday"
	default:
		result = "Some other day"
	}
	fmt.Println(result)

	// 15. Create loops with for statements
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 16. Write and Read local text files
	file, err := os.Create("./hello.txt")
	checkError(err)
	length, err := io.WriteString(file, fullName)
	fmt.Printf("Total characters %v\n", length)
	defer file.Close()
	// defer readFile("./hello.txt")

	// 17. Read file from Web
	res, err := http.Get(URL)
	checkError(err)
	fmt.Printf("Response Type - %T\n", res)
  defer res.Body.Close()

  bytes, err := ioutil.ReadAll(res.Body)
  checkError(err)
  fmt.Printf("Response Body - %v\n", string(bytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("File content -", string(data))
}
