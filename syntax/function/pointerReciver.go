package main
/*
		值接收者和指针接收者：
	* 使用场景：
	1. 需要改变内容必须使用指针接收者
	2. 结构过大也考虑使用指针接收者
	3. 一致性，如果有指针接收者，最好都使用指针接收者
 */
import "fmt"

type treeNode struct {
	value int
	left,right *treeNode
}

//值传递，与指针传递功能相同，值复制，函数体外无法看到值改变
func (node treeNode) print() {
	node.value = 300
	fmt.Print(node.value, " ")
}

//指针传递
func (node *treeNode) setValue(value int) {
	node.value = value
}

func main() {
	root := treeNode{value:3}

	root.setValue(100)
	root.print()
	fmt.Print(root.value)

	proot := &root
	proot.print()
	proot.setValue(200)
	proot.print()
}
