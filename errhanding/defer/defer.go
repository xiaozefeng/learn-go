package main

import (
	"bufio"
	"fmt"
	"imooc.com/learn-go/functional/fib"
	"os"
)

func main() {
	tryDefer()
	writeFile("fib.txt")
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok:= err.(*os.PathError); !ok{
			panic(err)
		}else {
			fmt.Printf("%s, %s, %s", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func tryDefer() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i," ")
	}
}
