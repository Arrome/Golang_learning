package main

import (
	"fmt"
	"strconv"
)

/**
	slice 代表变长的序列，每个元素都有相同类型。和数组相像，只是没有固定长度
		三部分构成： 指针、长度len、容量cap
	slice 不可以向前扩展，可以向后扩展。可以看到底层数组的后面的切片。不超过cap
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

	//slice 不可以向前扩展，可以向后扩展。可以看到底层数组的后面的切片。不超过cap
	fmt.Println(q[2:4])//q=[2 3] q[2:4]=[4,5]

	//apend方法，底层分配新的slice，1倍的cap增加
	num := []int{1,2,3}
	n := num[1:2]
	fmt.Println(n)
	fmt.Println("len(n):" + strconv.Itoa(len(n)) + ",cap(n):" +  strconv.Itoa(cap(n)))
	n1 := append(n, 4)
	n2 := append(n1, 5)
	n3 := append(n2, 6)
	n4 := append(n3, 7)
	fmt.Println(n1,n2,n3,n4)
	fmt.Println(strconv.Itoa(cap(n1)) + "," + strconv.Itoa(cap(n2)) + "," + strconv.Itoa(cap(n3)) + "," + strconv.Itoa(cap(n4)))

	//make函数返回一个slice
	slice := make([]int,3)
	fmt.Println(slice)

	//slice操作
	sliceappend := append(slice, 4)
	fmt.Print("after append:")
	fmt.Println(sliceappend)

	fmt.Println(len(sliceappend))
	fmt.Println(cap(sliceappend))

	//[]int 不表示数组，表示切片
	sliceCopy := make([]int,2,2)
	copy(sliceCopy,sliceappend)
	fmt.Println(sliceCopy)

	//删除，没有提供特定方法，使用切片方式
	fmt.Println(week)
	fmt.Println(append(week[:3],week[4:]...)) //删除中间元素
}
