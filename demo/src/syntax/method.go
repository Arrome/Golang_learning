package main

import "fmt"

type rec struct {
	width,height int
}

/*
	附加的参数会将该函数附加到这种类型上，相当于为这种类型定义了一个独占的方法
	附加参数叫做方法的接收器，建议类型的第一个字母
	当调用一个函数时，会对每个参数值进行拷贝，
	如果一个函数需要更新一个变量，或者函数中的其中一个参数太大希望避免默认拷贝，需要使用指针而不是对象来声明方法
*/
func (r *rec) area() int {
	return r.width * r.height
}

func main() {
	r := rec{width:10,height:5}
	fmt.Println("area: ",r.area())
}
