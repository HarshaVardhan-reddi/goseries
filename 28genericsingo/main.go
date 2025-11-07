// assign data type in the run time
// dont have to repeat the same code for different data typed inputs
// https://www.youtube.com/watch?v=WpTKqnfp5dY

package main

import (
	"fmt"

	// "golang.org/x/exp/constraints"
)


func AddInt(a int, b int) int{
	return a+b
}


func AddFloat(a float64, b float64) float64{
	return a+b
} // repeatitive code for float

// Generic defining way 1
func GenericAdd[T int | float64](a T, b T) T {
	return a+b
}

// Generic defining way 2
type Num interface {
	int | int8 | int16 | float32 | float64
}

func GenericAdd1[T Num](a T, b T) T {
	return a+b
}

// Generic defining way 3
// func GenericAdd2[T constraints.Ordered](a T, b T) T {
// 	return a+b
// }
// constraints.Ordered wraps all the primitive data types inside this
// works in go >= 1.24.0


func main(){
	fmt.Println("Welcome to generics in golang")
	result := AddInt(1,2)
	fmt.Printf("result: %v\n", result)

	floatres := AddFloat(2.3, 4.5)
	fmt.Printf("floatres: %v\n", floatres)

	// common function for everything
	genericresul := GenericAdd(1,2)
	fmt.Printf("genericresul: %v\n", genericresul)

	genericresul1 := GenericAdd(1.1,2.2)
	fmt.Printf("genericresul1: %v\n", genericresul1)
}