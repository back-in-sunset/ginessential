package main

import (
	"gin-essential/pkg/errors"
	"log"
	"os"
)

// C demo
func C() error {
	file, err := os.Open("abc")
	defer file.Close()
	if err != nil {
		err = errors.New("open file failed") // 对err进行包装，附带堆栈信息处理。
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
