package main

import "fmt"

func main(){
	fmt.Println("Welcome to loops in golang")

	days := []string{"monday","tuesday","wednesday","friday","saturday"}

	fmt.Println(days)

	// For loop
	// for i:=0; i<len(days); i++{
	// 	fmt.Println("Day:",days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i]) // i will be index
	// }

	for index, day := range days {
		fmt.Printf("Index is %v, and the value is %v \n",index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 5 {
			// break
			rougueValue ++;
			continue
		} else if rougueValue == 2{
			goto lco
		}
		fmt.Println("Value is: ",rougueValue)
		rougueValue ++
	}


	// Block for GOTO statements
	lco:
		fmt.Println("Jumping here!")
}