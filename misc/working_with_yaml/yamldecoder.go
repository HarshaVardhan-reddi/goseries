package workingwithyaml

import (
	"fmt"
	"strings"
	"github.com/goccy/go-yaml"
)

type SampleYaml struct{
	A int `yaml:"a"`
	B int `yaml:"b"`
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
}