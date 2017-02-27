package main
// cd src/runner && go install && cd ../../ && runner
import (
	"testing/quick"
	"reflect"
	"math/rand"
	"time"
	"fmt"
)

func main() {
	val, ok := quick.Value(reflect.TypeOf(1), rand.New(rand.NewSource(time.Now().Unix())))
	if ok {
		fmt.Println(val.Int())
	}

	val, ok = quick.Value(reflect.TypeOf(1.), rand.New(rand.NewSource(time.Now().Unix())))
	if ok {
		fmt.Println(val.Float())
	}

	val, ok = quick.Value(reflect.TypeOf(MyStruct{}), rand.New(rand.NewSource(time.Now().Unix())))
	if ok {
		fmt.Println(val.Interface().(MyStruct))
	}
}

type MyStruct struct {
	MyInt int
	MyString string
	MySlice []float32
}
