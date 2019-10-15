package main

func main() {


	/*
		defer 关键字，可以注册多个延迟调用，FIFO顺序在函数返回时执行。
		类似Java finaly，常用来保证一些资源最终一定能得到回收和释放

		defer后必须是函数或方法调用，不能是语句

		defer位置不当，可能导致panic，一般defer语句放在错误检查语句之后
		副作用：defer会推迟资源的释放，defer尽量不要放到循环语句里边（大函数内部defer拆分成小函数）
		defer相对于普通函数调用需要间接的数据结构支持，有一定的性能损耗

		defer陷阱：最好不要对有名返回值参数进行操作
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
	//os.Exit(1)
	a++
	return a
}