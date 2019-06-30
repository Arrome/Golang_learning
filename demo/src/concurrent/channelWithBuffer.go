package main

import (
	"fmt"
	"runtime"
)
/**
	关闭channel后：无法写数据，但是可以读数据
 */
func main() {
	ch := make(chan int,3) //存满即发生阻塞
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n",len(ch),cap(ch))

	go func() {
		for i := 0; i < 10; i++  {
			ch <- i
			fmt.Printf("子协程[%d]: len = %d,cap = %d\n",i,len(ch),cap(ch))

			if i == 7 {
				//不需要写数据时，关闭
				close(ch)
				break
			}
		}
	}()

	runtime.Gosched()
	//for i := 0; i < 10; i++ {
	//	if num,ok := <- ch; ok == true {
	//		fmt.Println("num = ",num)
	//	} else {
	//		break
	//	}
	//
	//}
	//range 遍历管道,检查到channel关闭，跳出
	for num := range ch {
		fmt.Println("num = ",num)
	}
}
