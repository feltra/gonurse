package main
// abstrair isso para nurse
// nurse manda o arquivo e as trigger parar testar

import (
	"fmt"
	"github.com/smallfish/simpleyaml"
	"io/ioutil"
	"os"
//	"strings"
)

func main() {
	filename := os.Args[1]
	//field := os.Args[2]
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	yaml, err := simpleyaml.NewYaml(source)
	if err != nil {
		panic(err)
	}
	a, _ := yaml.Array()
	for i := 0; i < len(a); i++ {
		//value, _ := yaml.GetIndex(i).Get("param").Get(field).String()
		m, _ := yaml.GetIndex(i).Map()
		//res := strings.EqualFold(value, check)
		//fmt.Printf("%s: %#v\n", field,x)
		for k,v := range m {
			fmt.Printf("%s - %s\n", k,v)}

	}
}
