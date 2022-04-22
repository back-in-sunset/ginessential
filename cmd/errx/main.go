package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(stack())
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}
