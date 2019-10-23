package main

import "fmt"

/**
	仅在defer调用中使用
	获取panic的值
	如果无法处理，可重新panic
 */

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()
}
func main() {
	
}
