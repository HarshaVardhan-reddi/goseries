package main

import "fmt"

// func methodname(*params) returntype {}, if void noting to mention example greeter function
func addTwoNumbers(a int, b int)int {
	var sum int = a+b
	return sum
}

// Entry point function to the program
// Anonymous functions exists in go
func main(){
	fmt.Println("Welcome to functions in go")
	greeter()
	result := addTwoNumbers(1,2)
	fmt.Println("output of the two numbers:",result)
	fmt.Println("Adding all the numbers passed to the function", proAdder(1,2,4,5,6))

	sum, message := multiReturns(1,2,3,4,4,4,4,10) // storing multivalued result
	fmt.Println("Message which we got: ", message, "\nand the sum is: ", sum)
}

func greeter(){
	fmt.Println("Namaste from golang")
}

// Pro functions
func proAdder(values ...int)int {
	sum := 0
	for _, element := range values {
		sum += element
	}
	return sum
}


// Pro functions
// Two values can be returned from a function
func multiReturns(values ...int)(int,string) {
	sum := 0
	for _, element := range values {
		sum += element
	}
	return sum, "Hello this is a string"
}
