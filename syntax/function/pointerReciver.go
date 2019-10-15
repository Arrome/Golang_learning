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

/*
	附加的参数会将该函数附加到这种类型上，相当于为这种类型定义了一个独占的方法
	附加参数叫做方法的接收器，建议类型的第一个字母
	当调用一个函数时，会对每个参数值进行拷贝，
	如果一个函数需要更新一个变量，或者函数中的其中一个参数太大希望避免默认拷贝，需要使用指针而不是对象来声明方法
*/
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
