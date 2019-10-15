package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)

	value := reflect.ValueOf(b)
	fmt.Println(value.Kind())

	v := value.Interface()
	if vv,ok := v.(string);ok {
		fmt.Println(vv)
	} else {
		fmt.Println("值类型不对")
	}
}

func main() {
	var a int = 200
	test(a)
}
