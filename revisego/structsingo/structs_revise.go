package structsingo

import "fmt"

type AdminUser struct {
	Name string
	hobbies []string
}

func PracticeStruct() {
	var abcObj AdminUser =  AdminUser{"abc", []string{"cricket"}}
	fmt.Println(abcObj.hobbies)
	fmt.Println(abcObj.Name)

	second := AdminUser{
		Name: "abc",
		hobbies: []string{"one","two","three"},
	}

	fmt.Println(second)
}