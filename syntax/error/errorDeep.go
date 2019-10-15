package main

import (
	"fmt"
)

/**
	错误处理问题：
	* 函数该如何返回错误，是用值，还是用特殊的错误类型
	* 如何检查被调用函数返回的错误，是判断错误值，还是用类型断言
	* 程序中每层代码在碰到错误的时候，是每层都处理，还是只用在最上层处理，如何做到优雅
	* 日志中的异常信息不够完整、缺少stack strace，不方便定位错误原因
 */


/*
	错误处理策略1：
		返回和检查错误值：通过特定值表示成功和不同错误，上层代码检查错误值，来判断被调用func的执行状态
 */
//处理不灵活，上层代码需要判断是否等于特定值，若修改，上层就要修改
//func strategy1() {
//	buf := make([]byte, 100)
//	n, err := r.Read(buf)
//	buf = buf[:n]
//	if err == io.EOF {
//		log.Fatal("read failed:", err)
//	}
//}

/*
	错误处理策略2：
		自定义错误类型：通过自定义错误类型来表示特定的错误，上层代码通过类型断言判断错误类型
    优势在于：可以将底层错误包起来一起返回给上层
*/

type MyError struct {
	Msg string
	File string
	Line int
}

func (err MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s",err.File,err.Line,err.Msg)
}

func dosomething() error {
	return MyError{"somethings happened","server.xml",22}
}

/*
	错误处理策略3：
		隐藏内部细节，降低耦合，上层只需要知道调用函数是否正常工作，不知道错误细节
	不返回具体错误内容，上层只通过 != nil 判断，可使用fmt.Errorf()打印具体信息调试
*/


func main() {
	err := dosomething()
	if _,ok := err.(MyError);ok{
		fmt.Println(err)
	}
}
