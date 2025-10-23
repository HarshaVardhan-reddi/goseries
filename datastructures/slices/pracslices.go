package slices

import(
	"fmt"
)

func Practice(){
	// defining slices
	var names = []string{"Harsha","reddy","Siri"}
	fmt.Println("printing names",names)

	makeNames := make([]string, 5)
	makeNames[0] = "HarshaMake"
	fmt.Println("make names", makeNames)

	for i:=0; i < 10; i++ {
		fmt.Println("Current index", i)
	}

	for index,value := range names {
		fmt.Println("index ", index,"value ",value)
	}
}