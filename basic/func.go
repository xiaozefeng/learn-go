package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	result, e := eval(1, 2, "x")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(result)
	}
	fmt.Println("---------")
	fmt.Println(
		div(13, 3),
	)
	fmt.Println("----------")
	fmt.Println(apply(3, 4, pow))
	fmt.Println("------------------")
	fmt.Println(apply(3, 4, func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}))
	fmt.Println("------------")
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println("-----------------")
	a, b = swap1(a, b)
	fmt.Println(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func eval(a, b int, ops string) (int, error) {
	switch ops {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsuppored operation: %s", ops)
	}
}

// 函数式编程
func apply(a, b int, op func(int, int) int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 13/3  4...1
func div(a, b int) (int, int) {
	return a / b, a % b
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int) {
	return b, a
}
