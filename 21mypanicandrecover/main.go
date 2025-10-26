package main

import "fmt"

func main(){
	fmt.Println("Welcome to pnaic tutorial")

	defer func ()  {
		r := recover()
	
		if(r != nil){
			// caching the error
			fmt.Println(r)
		}
	}()

	panic("Raising panic here!!!") // user only for unrecoverable errors
}