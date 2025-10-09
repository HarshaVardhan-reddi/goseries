package main

import(
	"fmt"
	rb "revise/basicproblems" // alias for an imported package
)

func main(){
	fmt.Println("Hello world in go")
	var x = 10
	fmt.Printf("printing x %d \n", x)

	var y int  = 10
	fmt.Printf("printing y %d\n", y)

	result := rb.IsEven(2)
	fmt.Println("Result is: ", result)
}