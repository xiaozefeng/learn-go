package basic

import "fmt"

func main() {
	var s []int // Zero value for slice is nil
	printSlice(s)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	printSlice(s)
	// creating slice
	fmt.Println("creating slice")
	s1 := make([]int, 16)
	printSlice(s1)
	s2 := make([]int, 10, 16)
	printSlice(s2)

	// copy slice
	fmt.Println("copy slice")
	copy(s1, s)
	printSlice(s1)
	// s1 = data=[0 1 2 3 4 5 6 7 8 9 0 0 0 0 0 0], length=16, cap=16
	// now delete 3
	fmt.Println("delete element from slice")
	s1 = append(s1[:3], s1[4:]...)
	fmt.Println(s1)

	fmt.Println("poping from front")
	front := s1[0]
	s4 := s1[1:]
	fmt.Println(front)
	fmt.Println(s4)

	fmt.Println("poping from end")
	end := s4[len(s4)-1]
	s5 := s4[:len(s4)-1]
	fmt.Println(end)
	fmt.Println(s5)

}

func printSlice(s []int) {
	fmt.Printf("data=%v, length=%d, cap=%d\n", s, len(s), cap(s))
}
