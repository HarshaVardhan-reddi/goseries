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
	validJson := json.Valid(bytedinput)
	if validJson {
		fmt.Println("Processing json as it is valid")
		response := map[string]interface{}{} // all the data types implemenmts an empty interface by default
		json.Unmarshal(bytedinput, &response)
		fmt.Println(response)
	} else {
		panic("Invalid json")
	}

	err := json.Unmarshal(bytedinput, &data)
	if err != nil {
		panic(err)
	}

	// fmt.Println(data)
	return data
}

// So in prod grade systems
// response := map[string]interface{}{} this will be used initially as a raw repsonse
// then the raw response will be validated and mapped to a defined struct