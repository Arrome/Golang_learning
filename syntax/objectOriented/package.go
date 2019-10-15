package main

/**
	* 一个目录下只能定义一个package包,每个目录下一个main文件
	* 为结构定义的方法放在同一个包下，可以是不同文件
    * 扩充系统类型..:定义别名，使用组合


	可以所有项目和第三方库都放在GOPATH下
*/

//import 别名命名方式
import (
	"bytes"
	l "log" //使用别名，避免重名
	. "fmt" //省略包名，直接使用函数
	_ "os" //只是引入该包，不调用，所有init()函数执行
	)

func main() {
	Println("hello")

	var (
		buf    bytes.Buffer
		logger = l.New(&buf, "logger: ", l.Lshortfile)
	)

	logger.Print("Hello, log file!")
}
