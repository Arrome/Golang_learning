/*
go install xxxdir：编译出来的可执行文件以其所在目录名命名，安装到src同级bin目录下(自动创建)，可执行文件依赖的各种package编译后，放在src同级pkg目录下

go build xxxGofile:默认情况可执行文件为go的文件名,生成到当前目录
go build : 将得到一个和当前文件夹同名的可执行文件,生成到当前目录

go run xxxGofile: 直接运行Go文件
 */

//包名可以和文件夹不同名，包名为 main 的包为应用程序的入口包，编译源码没有 main 包时，将无法编译输出可执行的文件。
package main

import "fmt"

func main() {
	fmt.Print("hello world" )
}
