// cd packt/learning_path_go/hello_world/ && go build -o hello_world && ./hello_world
package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	a int
	b string
}

func main() {
	pi := 3.141592
	fmt.Println("type:", reflect.TypeOf(pi))
	var a string
	a = "test"
	fmt.Println("type of string:", reflect.TypeOf(a))
	test := &Test{a: 1, b: "2"}
	fmt.Println("type of struct pointer:", reflect.TypeOf(test))
	test2 := Test{a: 1, b: "2"}
	fmt.Println("type of struct:", reflect.TypeOf(test2))
}