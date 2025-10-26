package kindaadverrors
// https://earthly.dev/blog/golang-errors/

import (
	"errors"
	"fmt"
)

func HandleErrors() error {
	err := errors.New("using errors new in error hanlding")
	if err != nil{
		fmt.Println(err)
	}
	return fmt.Errorf("using fmt for printing errors %d", 1)
}