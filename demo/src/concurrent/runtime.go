package main

import (
	"fmt"
	"runtime"
)

/**
	runtime.Gosched() 让出时间片先让别的协程执行，执行完之后再回到此协程
	runtime.Goexit() 终止所在的协程
	runtime.GOMAXPROCs() 设置可以并行计算的核数
 */
func main() {

	//go func() {
	//	for i := 0; i < 5; i++  {
	//		fmt.Println("go")
	//	}
	//}()
	//
	//for i := 0; i < 2; i++  {
	//	runtime.Gosched()
	//	fmt.Println("hello")
	//}
	//
	//go func() {
	//	fmt.Println("aaaaaaaa")
	//	test()
	//	fmt.Println("ddddddddd")
	//}()
	//
	//for {}

	n := runtime.GOMAXPROCS(2)
	fmt.Println(n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}

func test() {
	defer  fmt.Println("cccccccccc")
	//return
	runtime.Goexit()
	fmt.Println("dddddddddddd")
}
