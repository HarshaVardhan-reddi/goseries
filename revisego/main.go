package main

import(
	"fmt"
	"revise/basicproblems"
)

func main(){
	fmt.Println("Hello world in go")
	var x = 10
	fmt.Printf("printing x %d \n", x)

	var y int  = 10
	fmt.Printf("printing y %d\n", y)

	result := basicproblems.IsEven(2)
	fmt.Println("Result is: ", result)
}