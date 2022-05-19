package utils

import (
	"strconv"
	"time"
)

const (
	// DateTimeLayout 日期模版
	DateTimeLayout = "2006-01-02 03:04:05"
)

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

// ToTime timeStr to time
func (a S) ToTime(timeLayout string) (time.Time, error) {
	return ToTime(timeLayout, func() string {
		return string(a)
	})
}

// ToTime ..
func ToTime(timeLayout string, timeStrFn func() string) (time.Time, error) {
	return time.Parse(timeLayout, timeStrFn())
}
