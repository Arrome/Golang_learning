package main

import "fmt"

/**
 关键字 25 个：
	break default func interface select
	case  defer   go   map       struct
    chan  else    goto package   switch
    const fallthrough if range   type
    continue for  import return  var
  内建常量： true false iota nil
  内建类型： int int8 int16 int32 int64
			uint uint8 uint16 uint32 uint64 uintptr
			float32 float64 complex128 complex64
			bool byte rune string error
  内建函数：
	make len cap new append copy close
	delete complex real imag panic recover
 */


func main() {

	/*
		panic 主动抛出Go的runtime errors，recover捕获panic抛出的错误

		panic两种情况：
		1. 程序主动调用panic函数
			场景：
				1.程序遇到无法正常执行下去的错误，主动调用panic函数结束运行
				2.在调试时，主动调用实现快速退出，panic打印出堆栈更快定位错误
		2. 程序产生运行时错误，由运行时检测抛出

		panic过程：程序会从调用panic的函数位置或发生panic的位置立即返回，逐层向上执行函数的defer语句，然后打印函数调用堆栈，
				直到被recover捕获或运行到最外层函数退出
		recover捕获panic，阻止panic继续向上传递。只有在defer后面的函数体内被直接调用才能捕获panic终止异常，否则返回nil,继续向外传递
	 */
	//defer func() {
	//	println("函数体")
	//}()
	//
	//defer exception()
	//panic("test panic and recover")


	/*
		可以连续有多个panic被抛出，场景：延迟调用使用
		但是只有最后一个panic被捕获
	 */
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("first defer panic")
	}()

	defer func() {
		panic("second defer panic")
	}()

	panic("main body panic")
}

func exception() {
	println("调用函数")
	recover()
}

