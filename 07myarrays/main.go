package main

import "fmt"

func main(){
	fmt.Println("Welcome to arrays in golang")
	var fruitList[4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"
	fmt.Println("Fruit list is:", fruitList) // -> Fruit list is: [Apple Mango  Peach] | There will be space
	fmt.Println("Fruit list is:", len(fruitList)) // -> 4 (Just returns space registered for array, but not number of elememnts)

	var vegList = [5]string {"Potato", "Beans", "Mushroom"}
	fmt.Println("Veggie list is: ", vegList) // -> Veggie list is:  [Potato Beans Mushroom]
	fmt.Println("No of veggies", len(vegList)) // -> 5 (Just returns space registered for array)
}