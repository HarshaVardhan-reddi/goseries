package encoders

import (
	"encoding/json"
	"fmt"
)

type DataForEncode struct{
	Id int `json:"identifier"`
	Name string
	Price float32
	Tags []string
}

func EncodeCourses(data DataForEncode){
	fmt.Println("Encoding the data")
	result, err := json.MarshalIndent(data,"","\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("Result",string(result))
}