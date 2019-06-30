package main

import (
	"fmt"
	"time"
)

/**
	channel实现同步：
		管道通信方式
 */
//无缓冲，make(chan type) 等价于make(chan type,0)
var ch = make(chan int)

func main() {
	//len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n",len(ch),cap(ch))

	go person1()
	go person2()

	fmt.Println("\n结束")

	select {

	}

}

func Prenter(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}
func person1() {
	Prenter("hello")
	ch <- 666 //写数据,若未读取出，同时会被阻塞
}

func person2() {
	<- ch //从管道读数据，若无数据则阻塞
	Prenter("world")
}