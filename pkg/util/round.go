package util

import (
	"math"
	"strconv"
)

// S 字符串
type S string

// ToInt string转int
func (a *S) ToInt() int {
	return a.ToInt()
}

// ToFloat64 string to float64
func (a *S) ToFloat64() float64 {
	return a.ToFloat64()
}

// Round 保留两位小数
func Round(x float64) float64 {
	return math.Round(x*100) / 100
}

// FloatRoundToStr float 保留prec位小数，不足prec位补零 prec为-1保留原始精度 to string
func FloatRoundToStr(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}
