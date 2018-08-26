package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(n int) {
			for {
				fmt.Printf("Hello from goroutine %d \n", n)
			}
		}(i)
	}
	time.Sleep(time.Minute *2)
}
