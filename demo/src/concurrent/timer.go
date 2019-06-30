package main

import (
	"fmt"
	"time"
)

/**
	timer 定时到，只会响应一次
 */
func main() {
	<-time.After(2*time.Second)
	fmt.Println("当前时间：",time.Now())

}

func main2() {
	timer := time.NewTimer(1*time.Second)
	fmt.Println("当前时间：",time.Now())

	for {
		<- timer.C
		fmt.Println("timer")
	}
}

func main1() {
	timer := time.NewTimer(2*time.Second)
	fmt.Println("当前时间：",time.Now())

	//2s后，向timer.C写数据，有数据后就读
	t := <-timer.C //没有数据前阻塞
	fmt.Println("t = ",t)
}
