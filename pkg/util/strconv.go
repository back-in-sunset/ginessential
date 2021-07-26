package util

import "strconv"

// S 字符串
type S string

// ToInt string转int
func (a S) ToInt() int {
	i, _ := strconv.Atoi(string(a))
	return i
}

// ToFloat64 string to float64
func (a S) ToFloat64() float64 {
	f, _ := strconv.ParseFloat(string(a), 64)
	return f
}
