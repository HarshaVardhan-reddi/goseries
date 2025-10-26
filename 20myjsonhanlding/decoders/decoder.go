package decoders

import (
	"encoding/json"
	"fmt"
	"myjsonhanlding/encoder"
)

func ParseInputStringifiedJson(input string) encoders.DataForEncode {
	fmt.Println("Input string: ", input)
	// typecasting to byte
	bytedinput := []byte(input)
	fmt.Println(bytedinput)

	data := encoders.DataForEncode{}

	err := json.Unmarshal(bytedinput, &data)
	if err != nil {
		panic(err)
	}

	// fmt.Println(data)
	return data
}