package structsingo

import "fmt"


type User struct {
	Name string
	Age int
	Address map[string]string // Maps
	Hobbies []string // slices
}

func StructsInGo(){
	var userobj User = User{
		Name: "Harsha",
		Age: 12,
		Address: map[string]string{"county":"India"},
		Hobbies: []string{"Food","Music"},
	}
	fmt.Println(userobj)

	oneMoreObj := User{"Harsha",12,map[string]string{"country":"ind"},[]string{"Food"}}
	fmt.Println(oneMoreObj)

	// declaring custom type
	type number int
	x := number(12)
	fmt.Println(x)
	fmt.Printf("\n%T",x)
}