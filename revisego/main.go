package main

import(
	"fmt"
	rb "revise/basicproblems" // alias for an imported package
	maps "revise/mapsrevise"
	strcuts "revise/structsingo"
	"revise/pointers"
)

type User struct {
	name string
	age int
} // here type creates a new data type kinda

func main(){
	fmt.Println("Hello world in go")
	var x = 10
	fmt.Printf("printing x %d \n", x)

	var y int  = 10
	fmt.Printf("printing y %d\n", y)

	result := rb.IsEven(2)
	fmt.Println("Result is: ", result)

	user := User{"Harsha",27}
	fmt.Println("user", user)
	fmt.Printf("type of user %T", user)

	// var arr = [5]int{1,2,3,4,5} Array

	maps.ReviseMaps()
	strcuts.StructsInGo()

	strcuts.PracticeStruct()

	mapprac := map[string]int{"absb":1}
	for key,value := range(mapprac) {
		fmt.Println(key,value)
	}

	map1 := make(map[string]int, 0)
	map1["abc"] = 1

	for key,value := range(map1) {
		fmt.Println(key,value)
	}

	var intx int = 10

	address, val := pointers.FetchAddressOfInteger(&intx)
	fmt.Println("Address of the val: ", address, "value in the address: ", val)
}