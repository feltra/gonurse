package main
import "fmt"
import "strings"
import "io/ioutil"
import	"github.com/smallfish/simpleyaml"
type Library struct {
	files []string
}


func (l Library) CatalogFiles(path string) []string {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if strings.HasSuffix(file.Name(),".yaml") {
			l.files = append(l.files,file.Name())
		}
	}
	return l.files
}

type	Prescriptions []interface{}
// converter array map in map[string]map[string] ?
func (k Prescriptions) Learn(files []string) []interface{} {
	var all []interface{}
	for _, file := range files {
		//fmt.Println(file)
		source, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		yaml, err := simpleyaml.NewYaml(source)
		a, _ := yaml.Array()
		for i := 0; i < len(a); i++ {
			all = append(all,yaml.GetIndex(i))
		//	value, _ := yaml.GetIndex(i).Get("param").Map()
		//	x, _ := value["name"]
		//	fmt.Printf("%s: %#v\n", "name",x)
		}
		if err != nil {
			panic(err)
		}
	}
	return all
}

func main() {
	l := Library{}
	//for _,file := range l.CatalogFiles(".") {
	//	fmt.Println(file)
	//}
	k := Prescriptions{}
	all := k.Learn(l.CatalogFiles("."))
	for _, v := range all {
			fmt.Println(v)
			m,_ := v.(map[string]interface{})
			fmt.Println(m)
	}
}

//import (
//	"fmt"
//	"github.com/smallfish/simpleyaml"
//	"io/ioutil"
//	"os"
////	"strings"
//)
//
//func main() {
//	filename := os.Args[1]
//	field := os.Args[2]
//	source, err := ioutil.ReadFile(filename)
//	if err != nil {
//		panic(err)
//	}
//	yaml, err := simpleyaml.NewYaml(source)
//	if err != nil {
//		panic(err)
//	}
//	a, _ := yaml.Array()
//	for i := 0; i < len(a); i++ {
//		//value, _ := yaml.GetIndex(i).Get("param").Get(field).String()
//		value, _ := yaml.GetIndex(i).Get("param").Map()
//		x, _ := value[field]
//		//res := strings.EqualFold(value, check)
//		fmt.Printf("%s: %#v\n", field,x)
//	}
//
//}
