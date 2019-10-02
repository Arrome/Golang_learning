package main

import (
	"fmt"
	"math"
)

var(
	aa int
	bb string
	cc bool
)
func variableTypeDeduction() {
	//编译器自动决定类型
	var a,b,c,d = 2,3,"abc",true
	fmt.Println(a,b,c,d)
}

// 只能在函数里定义,函数外必须使用var声明
func variableShorter(){
	a,b := 3,"def"
	fmt.Print(a,b)
}

func consts() {
	const(
		filename string = "abc.txt"
		aa,bb = 3,4
	)


	var c int
	c = int(math.Sqrt(aa*aa + bb*bb))
	fmt.Println(filename,c)

}

// 无枚举关键字，以此方式实现枚举
func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,javascript,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}
func main() {
	variableTypeDeduction()

	consts()

	enums()
}
