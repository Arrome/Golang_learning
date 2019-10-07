package main

import "fmt"
/**
	不支持默认参数值
	不支持函数重载
	不支持函数嵌套（不支持命名函数嵌套，支持匿名函数嵌套）
 */

func plus(a int,b int) int {
	return a+b
}

func plusplus(a,b,c int) int {
	return a + b + c
}

//多个返回值
func vals() (int,int) {
	return 3,7
}

//可变参数
/**
	所有不定参数类型相同
	不定参数必须是函数最后一个参数
	不定参数相当于切片，对切片操作都适用不定参数
	切片可以作为参数给不定参数，切片后要加上...
	形参为不定参数和形参为切片的函数类型不同
 */
func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0
	for _,num := range nums {
		total += num
	}
	fmt.Println(total)
}

//闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//匿名函数
func add(a, b int) (sum int) {
	anonymous := func(x,y int) int {
		return x + y
	}
	return anonymous(a,b)
}

func main() {
	res := plus(1,2)
	fmt.Println("1+2 =",res)

	res = plusplus(1,2,3)
	fmt.Println("1+2+3 = ",res)

	a,b := vals()
	fmt.Println(a,b)

	_,c := vals()
	fmt.Println(c)

	sum(1,2,3)
	num := []int{1,2,3,4}
	sum(num ...)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())

	fmt.Println(add(3,4))


}
