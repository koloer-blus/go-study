package main

import "fmt"

func main() {
	var slice1 []int
	arr1 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1[%d]: %d, addr: %p \n", i, arr1[i], &arr1[i])
	}
	slice1 = arr1[1:4]

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%d]: %d, addr: %p \n", i, slice1[i], &slice1[i])
	}
	var slice2 []int = make([]int, 5, 6)

	fmt.Print(slice2, &slice2, cap(slice2), len(slice2))
	fmt.Print(append(slice2, 4555))
	slice := make([]int, 10)
	copy(slice, slice2)
	changeSlice(slice)
	fmt.Println(slice)
}

func changeSlice(slice []int) {
	slice[0] = 999
}
