package mapsrevise

import "fmt"

func ReviseMaps(){
	var userDetails map[string]int
	fmt.Println("\nInitialize map with no memory",userDetails)

	var declaringMap = map[string]int{"abc":123}
	fmt.Println("Initialize map with some values",declaringMap)

	var makeMap = make(map[string]int, 5)
	makeMap["bac"] = 12
	fmt.Println("Declare map with make keyword",makeMap)

	for key,value := range declaringMap {
		fmt.Println("Key is ",key,"value is", value)
	}

}