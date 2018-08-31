package main

import "fmt"

func main() {
	ch := make(chan int)
	var sl []int
	for i := 0; i < 10000; i++ {
		sl = append(sl, i)
	}

	go sum(sl[:len(sl)/4], ch)
	go sum(sl[:len(sl)/2], ch)
	go sum(sl[:(len(sl)/4)*3], ch)
	go sum(sl[(len(sl)/4)*3:], ch)
	a, b, c, d := <-ch, <-ch, <-ch, <-ch
	fmt.Println(a, b, c, d, a+b+c+d)
}
func sum(sl []int, ch chan int) {
	sum := 0
	for _, v := range sl {
		sum += v
	}
	ch <- sum
}
