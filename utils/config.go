package book
// abstrair isso para nurse
// nurse manda o arquivo e as trigger parar testar

import (
	"fmt"
	"github.com/smallfish/simpleyaml"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	field := os.Args[2]
	check := os.Args[3]
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
		value, _ := yaml.GetIndex(i).Get(field).String()
		res := strings.EqualFold(value, check)
		if res {
			value, err := yaml.GetIndex(i).Get("param").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %#v\n", field,value)
		}
	}

}
