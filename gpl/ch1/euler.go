package main

import (
	"fmt"
	"math"
)

func euler() {
	fmt.Println(complex(math.E, 1i*math.Pi) + 1)
}

func main() {
	euler()
}
