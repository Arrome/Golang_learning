package main

import "fmt"

//数组是值类型，传参数是通过值拷贝。
//Go语言不直接使用数组，而使用切片。
// [3]int 和 [5]int 被视作不同数据类型
func main() {
	//声明
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i,v := range a {
		fmt.Printf("%d %d\n",i,v)
	}

	for _,v := range a {
		fmt.Printf("%d\n",v)
	}

	//声明并初始化1
	var q [3]int = [3]int{1,2}
	fmt.Println(q[2])
	fmt.Println(q[1])

	//声明并初始化2,索引方式
	var iq = [3]int{2:1,1:2}
	fmt.Println("iq[0]:" ,iq[0])
	fmt.Println("iq[1]:" ,iq[1])
	fmt.Println("iq[2]:" ,iq[2])

	//[...]表示数组长度根据初始值个数计算
	q2 := [...]int{1,2,3}
	fmt.Printf("%T\n",q2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
