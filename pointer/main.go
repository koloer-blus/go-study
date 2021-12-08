package main

import "fmt"

func main() {
	var intPointer = 8
	var pointer = &intPointer
	fmt.Printf("var %d, pointer: %v", intPointer, pointer)
	*pointer = 200
	fmt.Print(intPointer)
}
