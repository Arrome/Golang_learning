package main

import "fmt"

/**
	slice 代表变长的序列，每个元素都有相同类型。和数组相像，只是没有固定长度
		三部分构成： 指针、长度len、容量cap
 */
func main() {
	//定义从第一个元素开始，那么第0个元素会被自动初始为空
	week := []string{1:"1",2:"2",3:"3",4:"4",5:"5",6:"6",7:"7"}
	q := week[2:4]
	fmt.Println(q)

	q2 := week[:4]
	fmt.Println(q2)

	q3 := week[1:]
	fmt.Println(q3)

	q4 := week[:]
	fmt.Println(q4)

	num := []int{1,2,3}
	n := num[1:3]
	fmt.Println(n)

	//make函数返回一个slice
	slice := make([]int,3)
	fmt.Println(slice)
	sliceappend := append(slice, 4)
	fmt.Println(sliceappend)
}
