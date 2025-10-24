// not much of diff with functions
// functions inside structs are called as methods
package main

import "fmt"

type User struct {
	Name string
	Email string
	Status int
	Age int
	oneAge int
}

func main(){
	user := User{
		Name: "harsha",
		Email: "harsha@learnyst.com",
		Status: 1,
		Age: 12,
		oneAge: 1,
	}

	user.GetStatus()
	user.setEmail("abc@gmail.com")
	fmt.Println("Origninal:", user.Email)

	user.setEmailInPlace("abc@gmail.com")
	fmt.Println("Origninal after passing by pointer:", user.Email)
}

func (u User) GetStatus(){
	fmt.Println("Is user acitve ? :",u.Status)
}

// copy of user is passed, so to manuplate origninal object pass pointer
func (u User) setEmail(email string) {
	u.Email = email
	fmt.Println("Email of this user is", u.Email)
}
// passing pointer
func (u *User) setEmailInPlace(email string) {
	u.Email = email
	fmt.Println("Email of this user is", u.Email)
}
