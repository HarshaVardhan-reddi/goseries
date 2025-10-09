package slices

import (
	"fmt"
	"sort"
)

func CreateAndTraverseSlice(){
	fmt.Println("Initializing a slice in go, under hood an array")

	var frutis = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("\n type of data %T", frutis) // []string - this is slices

	frutis = append(frutis, "Mango", "Bannana")
	fmt.Println("\nFruits:", frutis) // []string - this is slices

	frutis = append(frutis[1:],"jackfruit") // start from index 1, ignore all prev's
	fmt.Println(frutis)

	frutis = append(frutis[1:3],"grapes") // Stops at 3 and non inclusive
	fmt.Println(frutis)

	frutis = append(frutis[:3], "picka") // start at 0 Stops at 3 and non inclusive
	fmt.Println(frutis)

	// defining slice using make
	highScores := make([]int, 4)
	fmt.Println(highScores)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // index out of bound
	fmt.Println(highScores)

	highScores = append(highScores, 555,666,321) // Its going to re-allocate the memory - No index out of bound
	fmt.Println(highScores)

	sort.Ints(highScores) // available in slizes not in the array
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) // boolean

	fmt.Println("Deletion in slices...")
	// remove a vlaue in slice based on the index
	courses := []string{"reactjs", "js", "swift", "python","ruby"}
	fmt.Println("slice courses:",courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println("removed courses:",courses)

}