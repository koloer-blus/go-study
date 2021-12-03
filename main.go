package main

import "fmt"

const s = "GO-study"
const Pi = 3.1415926
const b string = "test _  b"
const (
	a = iota
	t = iota
	c
)

var Ga int = 99

const v int = 139

func GetGa() func() int {
	if Ga := 55; Ga > 60 {
		fmt.Println("GetGa if", Ga)
	}
	for Ga := 2; ; {
		fmt.Println("GetGa for", Ga)
	}
	return func() int {
		Ga += 1
		return Ga
	}

}

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
	temp := GetGa()

	fmt.Println("main", temp(), temp(), temp(), temp())
	v := 1
	{
		v := 2
		fmt.Println(v)

	}
	fmt.Println(v)

}
