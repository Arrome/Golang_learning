package main

import "fmt"

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

func main() {
	variableTypeDeduction()
}
