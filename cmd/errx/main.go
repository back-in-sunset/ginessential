package main

import (
	"gin-essential/pkg/errors"
	"log"
	"os"
)

// C demo
func C() error {
	_, err := os.Open("abc")
	if err != nil {
		err = errors.WithStack(err) // 对err进行包装，附带堆栈信息处理。
		return err
	}
	return nil
}

func main() {
	err := C()
	if err != nil {
		log.Printf("%+v", err)
	}
}
