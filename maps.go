package main

import "fmt"

func main() {

	m1 := map[string]string{
		"name": "jack",
		"age":  "18",
	}
	fmt.Println(m1)

	m2 := make(map[string]int)
	fmt.Println(m2)

	var m3 map[string]int
	fmt.Println(m3)

	s := m1["name"]
	fmt.Println(s)

	if name, ok := m1["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key doesn't exist")
	}

	fmt.Println("delete key 'name'")
	delete(m1, "name")
	if name, ok := m1["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key doesn't exist")
	}

	fmt.Println("------------")
	fmt.Println("查找最长不重复子串")
	fmt.Println(lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcahglablvhahrwah"))
	fmt.Println(lengthOfNonRepeatingSubStr("12"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))

}

func lengthOfNonRepeatingSubStr(s string) int  {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI , ok := lastOccurred[ch]; ok && lastI >= start{
			start = lastI +1
		}
		if i -start+1 > maxLength {
			maxLength  = i -start +1
		}
		lastOccurred [ch ]  = i
	}
	return maxLength

}