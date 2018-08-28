package main

import (
	"regexp"
	"fmt"
)

const text = `My email is peter@gmail.com@abc.com
email1 is abc@def.com
email2 is     kkk@qq.com
email3 is aaslgh@com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	// -1表示所有
	match := re.FindAllStringSubmatch(text, -1)
	for _, val := range match {
		fmt.Println(val)
	}
}
