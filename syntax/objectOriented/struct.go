package main

import "fmt"

type treeNode struct{
	value int
	left,right *treeNode
}
//工厂方法实现自定义构造的功能
func createNode(value int) *treeNode {
	return &treeNode{value:value} //返回局部变量，编译器决定此值分配在堆上，还是栈上
}

func main() {
	var root treeNode

	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)
	root.right.right = createNode(2)

	nodes := []treeNode{
		{value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)
}
