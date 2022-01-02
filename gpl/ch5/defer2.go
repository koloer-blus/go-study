package main

import (
	"os"
	"runtime"
)

func main() {
	defer printStack()
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
