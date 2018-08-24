package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "YES我要吃饭了" // utf-8
	fmt.Println(len(s))
	for _, value := range []byte(s) {
		fmt.Printf("%X ", value)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d, %X) ", i, ch)
	}

	fmt.Println()
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println()
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)  ", i, ch)
	}
}
