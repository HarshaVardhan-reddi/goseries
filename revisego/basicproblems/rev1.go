package basicproblems

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

const THIS_IS_CONST int8 = 10

type User struct{
	Name string
	Age int8
	Height float32
}

func RevisionOne(){
	// declaring vars 
	var x int = 10
	// var er error
	// reader := bufio.NewReader(os.Stdin)
	// inp, err := reader.ReadString('\n')
	// if err != nil{
	// 	panic(err)
	// }
	// if already declared, it can be used in the following way. If not declared then we can use := syntax for assinging type dynamically
	// x, er = strconv.Atoi(strings.Trim(inp,"\n"))
	// if er != nil{
	// 	panic(er)
	// }

	age := int8(x)
	fmt.Println("Printing the age of the user", age)
	fmt.Printf("Type of the age %T\n", age)
	fmt.Printf("Type of the x %T\n", x)

	switch(x%2 == 0){
	case true:
		fmt.Println("Even")
	case false:
		fmt.Println("odd")
	}

	exampleslice := make([]int,0)
	fmt.Println("Example slice")
	exampleslice = append(exampleslice, 12)
	for i := 0; i < 10; i++{
		exampleslice = append(exampleslice, i)
	}
	fmt.Printf("exampleslice: %v\n", exampleslice)

	fmt.Println("This the constant : ", THIS_IS_CONST)
	fmt.Printf("This the constant %T\n ", THIS_IS_CONST)

	printerWithType(exampleslice[0:1])
	exampleslice = append(exampleslice[1:], exampleslice[1:]...) // removing 12
	printerWithType(exampleslice)

	deconstructionOfAnSlice(exampleslice[0:2]...)
	deconstructionOfAnSlice(1,2,2,2,2,3,3,3,3,3) // this kinda params cannot ve passed to sliced params function
	slicedParams(exampleslice[0:2])
}

type genericvals interface{
	int | int8 | int64 | []int
}

func printerWithType[T genericvals](num T) {
	fmt.Println("Printing the value in generic", num)
}


func deconstructionOfAnSlice(x ...int){
	fmt.Print("deconstructed value",x)
	fmt.Printf("deconstricured %T\n",x)
}


func slicedParams(x []int){
	fmt.Print("deconstructed value",x)
	fmt.Printf("deconstricured %T\n",x)
}