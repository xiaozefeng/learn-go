package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for v := range ch{
		fmt.Printf("%d ", v)
	}
}
// Channels aren't like files, you don't need to close them
// Closing is only necessary
func fibonacci(n int, ch chan int) {
	x,y := 0,1
	for i := 0; i < n; i++ {
		ch <- x
		x ,y = y, x+y
	}
	close(ch)

}




