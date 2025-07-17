package main

import "fmt"

const LoginToken string = "ddhddhdkdhdhdkj" // Capital 'L' is a public variable. Will be accessed by any other file
// Basically login token is exported here and can be imported in any pof other files

func main() {
	var username string = "harsha"
	var age int = 12
	var isAdmin bool = true
	var x = 10
	var userHeight float32 = 6.7

	fmt.Println(username)
	fmt.Printf("%d \n",age)
	fmt.Printf("%d \n",x)
	fmt.Printf("Variable of type: %T \n", age)
	fmt.Printf("Is this user an admin %t", isAdmin)
	fmt.Printf("Height of the user is: %f \n", userHeight)

	// default values & aliases
	var anotherVariable int
	fmt.Println(anotherVariable) // -> 0
	fmt.Printf("Variable of type: %T \n", anotherVariable)

	// implicit type
	var webiste = "harsha.com"
	fmt.Println(webiste)

	// no var style
	noVarVariable := 1000 // works only inside method
	fmt.Println(noVarVariable)

	// const
	fmt.Printf("This is the login token %s", LoginToken)
	fmt.Printf("type of login token is %T", LoginToken )
}