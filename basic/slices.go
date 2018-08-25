package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:] = ", arr[2:])
	s1 := arr[2:6]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)
	fmt.Println("After Update")
	updateSlice(s1)
	fmt.Println("s1=", s1)
	fmt.Println("arr=", arr)
	fmt.Println()
	updateSlice(s2)
	fmt.Println("s2=", s2)
	fmt.Println("arr=", arr)
	fmt.Println("----------------")
	fmt.Println("ReSlice")
	fmt.Println(s1)
	s1 = s1[:3]
	fmt.Println(s1)
	s1 = s1[:2]
	fmt.Println(s1)

}

func updateSlice(s []int) {
	s[0] = 100
}
