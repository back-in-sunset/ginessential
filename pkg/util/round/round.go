package round

import (
	"math"
	"strconv"
)

// Round 保留两位小数
func Round(x float64) float64 {
	return math.Round(x*100) / 100
}

// FloatRoundToStr float 保留prec位小数 舍去，不足prec位补零 prec为-1保留原始精度 to string
func FloatRoundToStr(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}
