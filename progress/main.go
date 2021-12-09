package main

import (
	"fmt"
	"os"
)

func selectControl() {
	//common ok
	if fileStream, err := os.Open("../docs/1-1.md"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Success Open the file!!!")
		fmt.Println(fileStream)
	}

	switch vari := 100; {
	case vari >= 100:
		fmt.Println("1")
		fallthrough
	case vari < 100:
		fmt.Println("2")
	}

	var temp interface{}
	var tempFloat = 3.1415926
	temp = tempFloat
	switch typeV1 := temp.(type) {
	case float32:
		fmt.Println("float32", typeV1)
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("other")
	}

	tempName := "Hello World"
	for key, value := range tempName {
		fmt.Printf("key=%d, value=%c, type=%T\n", key, value, value)
	}
}

func gotoFunc() {
	temp1 := 0
GOTO:
	fmt.Println(temp1)
	temp1++
	goto GOTO
}

func main() {
	selectControl()
}
