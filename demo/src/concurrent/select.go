package main

import (
	"fmt"
	"time"
)

/**
	每个case语句里必须是一个IO操作
 */
func main() {
	ch := make(chan int) //控制数字通信
	quit := make(chan bool) //控制程序结束



	go func() {
		for {
			select {
				case num := <-ch :
					fmt.Println("num = ",num)
					case <-time.After(3*time.Second):
						fmt.Println("超时")
						quit <- true
			}
		}
	}()

	<- quit
	fmt.Println("结束")
}

func main0() {
	ch := make(chan int) //控制数字通信
	quit := make(chan bool) //控制程序结束



	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println("num = ", num)
		}
		quit <-true
	}()

	fibonacci(ch,quit)
}

func fibonacci(ch chan<- int,quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
			case ch <- x:
				x, y = y, x+y
			case flag := <-quit:
				fmt.Println("flag = ",flag)
				return
		}
	}
}
