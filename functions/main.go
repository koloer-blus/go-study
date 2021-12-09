package main

import "fmt"

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
}

func computeValue(v1, v2 int) (sum int, minus int) {
	sum = v1 + v2
	minus = v1 - v2
	return sum, minus
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
