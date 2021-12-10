package main

import (
	"fmt"
	"os"
)

/**
* func funcName(...params) (returns) {

}
*/
func funcBase() {

}

func main() {
	funcBase()
	sum, minus := computeValue(4, 3)
	fmt.Println(sum, minus)
	mapArgs(1, "test", 3, 45.5353, false, []int{1, 2, 3, 4})
	fmt.Print(fibnaci(16))

	sum1 := func(v1, v2 int) int {
		return v1 + v2
	}(1, 2)
	fmt.Println(sum1)
	fmt.Println(addCount()(5))
	fmt.Println("defer")
	deferAction()

	fmt.Println(
		useFuncArgs(
			func(v1, v2 int) int { return v1 + v2 },
			func(v1, v2 int) int { return v1 - v2 },
		))
}

func deferAction() {
	handle, err := os.Open("../docs/1-1.md")
	fmt.Println(err)
	//async to run
	defer handle.Close()
}

func computeValue(v1, v2 int) (sum int, minus int) {
	sum = v1 + v2
	minus = v1 - v2
	return sum, minus
}

func useFuncArgs(addCom func(int, int) int, minuesCom func(int, int) int) (int, int) {
	return addCom(3, 5), minuesCom(6, 4)
}

func mapArgs(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Println("int:", v)
		case string:
			fmt.Println("string", v)
		case float32:
			fmt.Println("float32", v)
		default:
			fmt.Println("other")
		}
	}
}

func fibnaci(pot int64) int64 {
	if pot == 1 || pot == 2 {
		return 1
	} else {
		return fibnaci(pot-1) + fibnaci(pot-2)
	}
}

func addCount() func(int64) int64 {
	var step int64 = 0
	return func(_step int64) int64 {
		step += _step
		return step
	}
}
