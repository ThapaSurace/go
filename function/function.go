package main

import "fmt"

func main() {
	//getting two returned value
	name, age := details("Surace", 10)
	fmt.Printf("Name: %s, Age: %v\n", name, age)
	fmt.Println("------------------------------------")
	//calling add function
	a := 10
	b := 20
	res := add(a, b)
	fmt.Printf("The sum of %v and %v is %v\n", a, b, res)
	fmt.Println("------------------------------------")
	//ignoring return value using _
	num1, _ := numbers()
	fmt.Printf("The value of %v is\n", num1)
}

//returns two value of string and integer
func details(name string, age int) (string, int) {
	return name, age
}

//returns single int result
func add(x, y int) int {
	return x + y
}

func numbers() (int, int) {
	return 10, 20
}
