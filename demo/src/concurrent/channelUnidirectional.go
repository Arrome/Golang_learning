package main

import "fmt"

/**
	channel默认是双向的。
		ch := make(chan int)
		隐式转换为单向：(单向不可转换为双向)
			var writeCh chan<- int = ch
			var readCh <-chan int = ch
 */
func main() {
	ch := make(chan int)

	//生产者，写入channel
	go producer(ch)

	//消费者，从channel读取
	consumer(ch)
}

func producer(out chan<- int){
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ",num)
	}
}
