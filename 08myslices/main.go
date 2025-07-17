package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")
	var fruitlist = []string{ "Apple","Tomato", "Peach" }
	fmt.Printf("fruit list is of type: %T \n", fruitlist) // -> fruit list is of type: []string (this is acually slices)
	
	fruitlist = append(fruitlist, "PineApple", "Bananna")

	fmt.Println(fruitlist) // -> [Apple Tomato Peach PineApple Bananna]

	fruitlist = append(fruitlist[1:]) // Array slicing takes all the values after index 1 to rest

	fmt.Println(fruitlist) // -> [Tomato Peach PineApple Bananna]
	fmt.Println(len(fruitlist)) // => 4


	// [Tomato Peach PineApple Bananna]
	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	// using make keyword to create slices
	highScores := make([]int, 4)

	highScores[0] = 13
	highScores[1] = 23
	highScores[2] = 31
	highScores[3] = 41

	// highScores[4] = 777 // -> Out of the bound

	highScores = append(highScores, 555,104)

	fmt.Println(highScores) // -> [13 23 31 41 555 104]

	sort.Ints(highScores)
	fmt.Println(highScores) // -> 13 23 31 41 104 555]

	fmt.Println(sort.IntsAreSorted(highScores)) // -> true

	// removing elements from slice based on index ->>> https://youtu.be/931nR5TGCAk?si=d9kSSuLQ2cpbsMBs

	var courses = []string{"reactjs","javascript","swift","c","python","ruby"}
	fmt.Println(courses) // -> [reactjs javascript swift c python ruby]

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // [a:b] always b value is non inclusive. B will never be included in the list

	fmt.Println(courses) // [reactjs javascript c python ruby]

}