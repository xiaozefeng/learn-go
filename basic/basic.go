package basic

import (
	"fmt"
	"strconv"
)

func main() {
	variables()
}

func variables() {
	var a, b int = 1, 2
	var s1, s2 string = "a", "b"
	fmt.Println(a, b)
	fmt.Println(s1, s2)
	s := strconv.Itoa(1)
	fmt.Println(s)
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	bb := []byte(s)
	fmt.Println(bb)

	fmt.Println("------------")
	enums()
}

func enums() {
	const (
		golang = iota
		java
		python
		javascript
	)
	fmt.Println(golang, java, python, javascript)
}
