package workingwithyaml

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

type SampleYaml struct{
	A int `yaml:"a"`
	B int `yaml:"b"`
}

type ExampleYaml struct{
	A int `yaml:"a"`
	B int `yaml:"b"`
	C int `yaml:"c"`
}

type ExampleParent struct{
	Numbers ExampleYaml `yaml:"numbers"`
}

func DecodeSampleYaml() {
	sample := SampleYaml{}
	// there should be no spaces in YAML. So removed intendation
	yml := `
---
a: 1
b: 2
`
	yml = strings.TrimSpace(yml)

	err := yaml.Unmarshal([]byte(yml), &sample)
	if err != nil {
		panic(err)
	}
	fmt.Println("Unmarshalled YAML", sample)

	// converting unmarshlled yaml to json
	fmt.Println("Printing json")
	writter := os.Stdout
	err = json.NewEncoder(writter).Encode(sample)

	if err != nil{
		panic(err)
	}
}


func DecodeYamlFile(){
	file, err := os.OpenFile("/Users/harshavardhanreddy/goseries/misc/working_with_yaml/example.yml",0,os.FileMode(os.O_RDONLY))
	if err != nil{
		panic(err)
	}
	bytedoutput, errread := io.ReadAll(file)
	if errread != nil{
		panic(errread)
	}
	examplefile := ExampleParent{}
	if err := yaml.Unmarshal(bytedoutput, &examplefile); err != nil{
		panic(err)
	}
	fmt.Printf("examplefile: %v\n", examplefile)
}