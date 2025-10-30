package main

import "fmt"

// interfaces in go are implemented implicitly
// any struct having exact same signature of func() will be implemented by the interface automatically
// in the below example, Human interface has the method which is exactly implemented in the structs

// https://medium.com/@mbinjamil/using-interfaces-in-go-the-right-way-99384bc69d39

type Person struct{
	Name string
	Age int
}

type Admin struct{
	Name string
	Age int
}

type Customer struct{
	Name string
	Age int
}

type Human interface{
	DisplayName()
}

func main(){
	fmt.Println("Welcome to interfaces in golang")
	person := Person{
		Name: "Harsha",
		Age: 12,
	}

	customer := Customer{
		Name: "Jit",
		Age: 15,
	}

	admin := Admin{
		Name: "Suri",
		Age: 20,
	}

	humans := []Human{person,customer,admin}
	for _, human := range(humans){
		fmt.Println(human)
		human.DisplayName()
	}

	// p := Printer()
	// p.DisplayName
}

func(p Person) DisplayName(){
	fmt.Println("Hey person here!", p.Name)
}

func(a Admin) DisplayName(){
	fmt.Println("Hey Admin here!", a.Name)
}

func(c Customer) DisplayName(){
	fmt.Println("Hey customer here!", c.Name)
}