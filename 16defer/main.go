package main

import "fmt"

// import "fmt"

func main(){
	defer fmt.Println("Hellow world") // whatever is mentioned here, will be executed at the end of the function
	defer fmt.Println("One")
	defer fmt.Println("Two") // for multiple defers, LIFO will be followed, stacked up kinda
	fmt.Println("Hellow this is defer")
	myDefer()


}


func myDefer(){
	for i:=0; i < 10; i++ {
		defer fmt.Println(i)
	}
}