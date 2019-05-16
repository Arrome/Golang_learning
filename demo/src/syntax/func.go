package main

import "fmt"

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
}
