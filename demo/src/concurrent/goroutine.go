package main

import (
	"fmt"
	"time"
)
/**
	goroutine --协程，main函数即一个主协程
		比线程更小，Go语言内部实现goroutine间内存共享。执行goroutine只需要极少栈内存(大概4~5KB，根据相应数据伸缩)

	主协程退出，子协程自动退出
 */
func main() {

	go newTask() //新建一个协程,新建一个任务，放在前面

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}
}

func newTask() {
	for {
		fmt.Println("this is a new task")
		time.Sleep(time.Second)
	}
}