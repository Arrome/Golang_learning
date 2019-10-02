package main

import "fmt"

/**
	无序的key/value对集合 map[key]var
	内置非线程安全，可使用标准包sync中map
 */

func main() {
	//直接定义
	names := map[string]string {
		"gougou": "jinru",
		"maomao": "liyu",
	}
	fmt.Println(names["gougou"])

	//make函数创建一个map
	ages := make(map[string]int)
	ages["alice"] = 13
	ages["charlie"] = 32
	fmt.Println(ages["alice"])

	//删除map值
	delete(ages,"alice")
	fmt.Println(len(ages))

	//禁止对map元素取址（map可能随着元素数量的增长而重新分配更大的内存空间，导致之前地址无效）
	//_ = &ages["charlie"]

	//map迭代顺序随机，
	for name,age := range ages {
		fmt.Printf("%s\t%d\n",name,age)
	}
	for name := range ages {
		fmt.Println(name)
	}

	for i,c := range "go" {
		fmt.Println(i,c) // 1 111
	}
}
