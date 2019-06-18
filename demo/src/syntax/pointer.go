package main

import "fmt"

func main() {
	x := 1
	p := &x  // & 取地址

	//Go不支持指针运算，语言层面禁止。垃圾回收机制，指针运算会造成很多不便
	//p++

	fmt.Println(p)
	fmt.Println(*p)  // * 表示指针取值
	*p = 2
	fmt.Println(x)

	var m, n int
	fmt.Println(&m == &n, &m == nil)

	var p2 = f()
	fmt.Println(p2)

	//new函数创建匿名变量并初始化零值，返回变量地址，返回指针类型为*T
	p3 := new(int)
	fmt.Println(*p3)
}

func f() *int {  // 指针类型 *T
	v := 1
	return &v //Go编译器使用“栈逃逸”机制将局部变量的空间分配在堆上
}
