package main

import "fmt"

func main(){
	fmt.Println("Welcome to maps in golang")

	programmingLangs := make(map[string]string)
	programmingLangs["js"] = "JavaScript"
	programmingLangs["rb"] = "ruby"
	programmingLangs["py"] = "python"

	fmt.Println("list of all langs: ", programmingLangs) // -> map[js:JavaScript py:python rb:ruby]
	fmt.Println("JS has the value as: ", programmingLangs["js"]) // -> JavaScript

	delete(programmingLangs,"js")
	fmt.Println("list of all langs: ", programmingLangs) // map[py:python rb:ruby]
	fmt.Println("JS has the value as: ", programmingLangs["js"]) // empty
	fmt.Printf("JS has value of type %T \n",programmingLangs["js"]) // String as declared as string

	// traversing a map in golang
	for key, value := range programmingLangs{
		fmt.Println("Key is: ", key, "Value is: ", value)
	}

	for _, value := range programmingLangs{
		fmt.Printf("value of programming langs: %v \n",value) // %v can be used to print most types (integers, strings, structs, slices, maps, etc.) using their default format.
	}
}