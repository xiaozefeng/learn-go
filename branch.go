package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strconv"
)

func main() {

	filename := "a.txt"
	if c, err := ioutil.ReadFile(filename); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", c)
	}

	fmt.Println(grade(100))
	fmt.Println(grade(90))
	fmt.Println(grade(80))

	fmt.Println("------------------")
	fmt.Println(
		convertToBin(1),
		convertToBin(100),
		convertToBin(13),
		convertToBin(797),
	)


}

func convertToBin(n int) string {
	result := ""
	for ; n >0 ; n /=2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/**
 计算成绩
 */
func grade(score int) string {
	switch {
	case score < 60:
		return "F"
	case score < 70:
		return "E"
	case score < 80:
		return "C"
	case score < 90:
		return "B"
	case score <= 100:
		return "A"
	default:
		panic(fmt.Sprintf("Wrong scorr: %d", score))

	}
}
