package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	quit := make(chan int)

	go func() {
		// 一直取，没有东西放进ch 之前一直阻塞
		for i := 0; i < 20; i++ {
			fmt.Printf("%d ", <-ch)
		}
		// 放一个元素进quit
		quit <- 0
	}()

	fibonacci(ch, quit)
}
func fibonacci(ch chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// 当有人一直取时，一直执行此逻辑， for i := 0; i < 20; i++ { 这里取完了，
		// 元素就放不进去了， 然后放了一个元素进quit channel 这个时候执行 <- quit
		// 整个程序就退出了
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
