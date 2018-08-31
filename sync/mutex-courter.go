package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	c := SafeCounter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 1000; i++ {
			c.inc("some_key")
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			c.inc("some_key")
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(c.value("some_key"))

}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) inc(key string) {
	c.mux.Lock()
	c.v[key] ++
	defer c.mux.Unlock()
}

func (c *SafeCounter) value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
