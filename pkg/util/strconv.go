package util

import (
	"fmt"
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
	return time.Parse(timeLayout, string(a))
}

// ToMidnightTime 午夜零点
func (a S) ToMidnightTime() time.Time {
	t, _ := time.Parse(DateTimeLayout, fmt.Sprintf("%s 00:00:00", a))
	return t
}
