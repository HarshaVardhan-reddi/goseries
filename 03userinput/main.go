package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welome := "**welcome to user input**"
	fmt.Println(welome)
	// Comma ok syntax
	// declare a reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our pizza:")
	
	// comma ok // error ok
	input, err := reader.ReadString('\n')
	if(err != nil){
		fmt.Println("Hey its crashed here at line 20", err)
	}

	fmt.Println("Thanks for the rating:",input)
	fmt.Printf("type of the rating is %T", input)
}