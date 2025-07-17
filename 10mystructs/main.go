package main

import "fmt"

func main(){
	fmt.Println("Welcome to structs in golang")
	var x int = 123
	fmt.Println(x)
	// No inheritance in go lang; No super class or parent
	user := User{"Harsha","harsha@dev.go",true,16}
	fmt.Println(user) // {Harsha harsha@dev.go true 16}
	fmt.Println(user.Name) // Harsha

	// Another way of creating objects for structs
	var user1 User = User{"Harsha12","harsha12@dev.go",true,1612}
	fmt.Println(user1) // -> {Harsha12 harsha12@dev.go true 1612}
	fmt.Printf("user details are: %+v\n",user)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}