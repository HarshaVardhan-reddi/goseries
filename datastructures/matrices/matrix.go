package matrices

import "fmt"

func PrintMatrix() {
	fmt.Println("Welcome to matrix in golang")
	matrix := []([]int){}
	row1 := []int{1,2,3,4,5}
	row2 := []int{6,7,8,9,10}

	matrix = append(matrix, row1, row2)

	fmt.Println(matrix)
	fmt.Println("Printing matrix")
	for _,val := range(matrix){
		fmt.Print("\n")
		for _,val2 := range(val){
			fmt.Print(val2,"\t")
		}
	}
}