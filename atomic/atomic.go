package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var a atomicInt
	a.increase()
	go func() {a.increase()}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increase() {
	fmt.Println("safe increase")
	func(){
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value ++
	}()
}
func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
