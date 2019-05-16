package main

import "fmt"

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

	//声明并初始化
	var q [3]int = [3]int{1,2}
	fmt.Println(q[2])
	fmt.Println(q[1])

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
