package main

import "fmt"

func main(){
	fmt.Println("welcome to pointers")
	var one int = 2
	var ptr *int = &one
	fmt.Println("Value stored in pointer:", ptr)
	fmt.Println("Deref ptr:", *ptr)
}