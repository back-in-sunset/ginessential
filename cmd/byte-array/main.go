package main

import (
	"fmt"
)

func main() {
	var src = make([]byte, size)
	given(src)
	fmt.Println(src)
}

const size = 40

func given(src []byte) {
	var a byte = 'a'
	for i := 0; i < size; i++ {
		src[0] = a + byte(i)
		src = src[1:]
	}
}
