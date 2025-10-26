package main

import(
	"fmt"
	encoders "myjsonhanlding/encoder"
)

type course struct{
	Id int
	Name string
	Price float32
	Tags []string
}

func main(){
	crx := constructData()
	
	data := encoders.DataForEncode{
		Name: crx.Name,
		Id: crx.Id,
		Price: crx.Price,
		Tags: crx.Tags,
	}

	fmt.Println(crx)
	encoders.EncodeCourses(data)
}

func constructData() course {
	result := course{
		Id: 1,
		Name: "Go with golang",
		Price: 12.3,
		Tags: []string{"Excellent", "Super"},
	}
	return result
}