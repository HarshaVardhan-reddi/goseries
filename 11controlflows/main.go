package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to if else in golang")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	loginCount, error := strconv.ParseInt(strings.TrimSpace(input),10,64)
	if(error != nil){
		panic(error)
	}else if(loginCount%2 == 0){
		fmt.Printf("Login count is an even number : %v\n", loginCount)
	} else if(loginCount > 1){
		fmt.Printf("Login count is greather than 1 and value is : %v\n", loginCount)
	} else {
		fmt.Printf("Login count of user is: %v\n",loginCount)
		fmt.Printf("Type of Login count of user is: %T\n",loginCount)
	}
	
	// Initializing a variable & immediately checking the condition
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else{
		fmt.Println("num is not less than 10")
	}
}