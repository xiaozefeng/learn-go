package main

import (
	"fmt"
	"imooc.com/learn-go/retriever/mock"
	"imooc.com/learn-go/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}
func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is imooc.com"}
	inspect(r)

	r = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspect(r)

	// type assertion
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever ")
	}
	fmt.Println(download(r))

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
