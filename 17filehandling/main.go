package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("Welcome to golang file hanlding")
	// creating and writing into the file
	file, err := os.Create("gofilehanlding.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, "This is testing for lang")
	if err != nil {
		panic(err)
	}

	fmt.Println("Length of the content inside the file is: ", length)
	
	// reading from the file
	byteoutput, err := os.ReadFile("/Users/harshavardhanreddy/goseries/17filehandling/gofilehanlding.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Output of file read will be in byte format by default: ", byteoutput)
	fmt.Println("Converting byte to string", (string)(byteoutput))
}