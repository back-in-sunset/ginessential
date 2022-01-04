package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// timeNano()
	// parseTime()
}

const size = 40

func given(src []byte) {
	var a byte = 'a'
	for i := 0; i < size; i++ {
		src[0] = a + byte(i)
		src = src[1:]
	}
}

func group(num int) int {
	if num/10 >= 0 && num%10 > 0 {
		return num/10 + 1
	}
	return num / 10
}

type s struct {
	Ts time.Time `json:"ts"`
}

var os = `{"ts":1639115373898}`

func timeNano() {
	var s s
	buffer := bytes.NewBuffer([]byte(os))
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(&s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", s)
}

func parseTime() {
	var os = `{"ts":"2019-10-12T07:20:50Z"}`

	var s s
	err := json.Unmarshal([]byte(os), &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", s)
}

var data = `{"name":"jerry","age":"35","sex":"male"}`

var cars = []string{"toyota", "p7", "nisson"}

var params = struct {
	testString      string
	testInt         int
	testIntArray    []int
	testStringArray []string
}{
	testString:      "string",
	testInt:         11,
	testIntArray:    []int{-1, -2, -3},
	testStringArray: []string{"string1", "string2"},
}
