package main

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
import "os"
func main() {


	/*
		defer 关键字，可以注册多个延迟调用，FIFO顺序在函数返回时执行。
		类似Java finaly，常用来保证一些资源最终一定能得到回收和释放

		defer后必须是函数或方法调用，不能是语句
	 */
	defer func() {
		println("first")
	}()

	defer func() {
		println("second")
	}()

	println("funciton body")

	println(fdefer())
}


func fdefer() int {
	a := 0
	defer func(i int) {
		println("defer i=",i)
	}(a) //值拷贝方式传递，结果不受后续影响

	//退出进程，defer不再执行
	os.Exit(1)
	a++
	return a
}