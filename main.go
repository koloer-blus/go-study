package main

import "fmt"

const s = "GO-study"

func main() {
	var (
		a int = 1
		b bool
		c string
	)
	x := [3]int{1, 2, 3}
	fmt.Print("Hello world!!!\n")
	fmt.Print(a, b, c)
	for i, r := range s {
		fmt.Printf("%#U, %d \n", r, i)
	}

	for i, r := range x {
		fmt.Printf("%#U, %d \n", r, i)
	}
}
