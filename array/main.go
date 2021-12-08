package main

import "fmt"

func changeArr(arr [10]int) {
	arr[0] = 100
}

func changeArrByPoint(arr *[10]int) {
	(*arr)[0] = 120
}

func arrFunc() {
	arr := [10]int{}
	changeArr(arr)
	fmt.Println(arr)
	changeArrByPoint(&arr)
	fmt.Println(arr)
}

func main() {
	var array1 [10]int = [10]int{1, 3, 4, 67, 5}
	array1[0] = 100
	fmt.Println(array1, len(array1))

	for i := 0; i < len(array1); i++ {
		fmt.Printf("array1[%d] = %d, %p \n", i, array1[1], &array1[i])
	}

	arr2 := [...]int{1, 2, 3, 4}
	fmt.Println(arr2)

	arr3 := [...]int{200: 1}
	fmt.Println(len((arr3)))

	var arr4 [4][6]int = [4][6]int{
		{2, 34, 5, 3, 4, 54},
		{3, 455, 76, 4, 6},
		3: {9},
	}
	fmt.Println(arr4)

	arrFunc()
}
