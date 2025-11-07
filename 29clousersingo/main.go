package main

import "fmt"

func doubleElementsInSlices(slc []int, doubler func(int) int) []int { // callback function
	var tempslc []int
	for _, val := range slc{
		tempslc = append(tempslc, doubler(val))
	}
	return tempslc
}

// returning a closure(A function basically which can wrap varibles around it)
func returnAdderWithExtras() func(int,int) int{
	extraAdder := 1
	adder := func(a,b int)int{
		return a+b + extraAdder
	} // this is clouser which wraps the variables around it into the execution stack. Here its wrapping extraAdder
	return adder
}

func main() {
	fmt.Println("Welcome to closures in golang")
	result := doubleElementsInSlices([]int{1,2,3,4,5}, func(x int) int{
		return x*2
	}) // technically its a closures, but its a callback. 
	fmt.Printf("result: %v\n", result)

	// closures
	adder := returnAdderWithExtras()
	adder_with_extras := adder(3,5)
	fmt.Printf("adder_with_extras: %v\n", adder_with_extras)
}