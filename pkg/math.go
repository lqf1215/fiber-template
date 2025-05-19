package pkg

import "github.com/shopspring/decimal"

// DecimalRoundFixNum 四舍五入
func DecimalRoundFixNum(val float64, num int32) float64 {
	v, _ := decimal.NewFromFloat(val).Round(num).Float64()
	return v
}

// DecimalMultiplyRoundFixNum 相乘
func DecimalMultiplyRoundFixNum(val1 float64, val2 float64, num int32) float64 {
	fromFloat1 := decimal.NewFromFloat(val1)
	fromFloat2 := decimal.NewFromFloat(val2)
	v, _ := fromFloat1.Mul(fromFloat2).Round(num).Float64()
	return v
}

// DecimalAddRoundFixNum 相加
func DecimalAddRoundFixNum(val1 float64, val2 float64, num int32) float64 {
	fromFloat1 := decimal.NewFromFloat(val1)
	fromFloat2 := decimal.NewFromFloat(val2)
	v, _ := fromFloat1.Add(fromFloat2).Round(num).Float64()
	return v
}

// DecimalSubtractRoundFixNum 相减
func DecimalSubtractRoundFixNum(val1 float64, val2 float64, num int32) float64 {
	fromFloat1 := decimal.NewFromFloat(val1)
	fromFloat2 := decimal.NewFromFloat(val2)
	v, _ := fromFloat1.Sub(fromFloat2).Round(num).Float64()
	return v
}

// DecimalDivideRoundFixNum 相除
func DecimalDivideRoundFixNum(val1 float64, val2 float64, num int32) float64 {
	fromFloat1 := decimal.NewFromFloat(val1)
	fromFloat2 := decimal.NewFromFloat(val2)
	v, _ := fromFloat1.Div(fromFloat2).Round(num).Float64()
	return v
}

// DecimalCompareTo 比较大小
func DecimalCompareTo(val1 float64, val2 float64) int {
	fromFloat1 := decimal.NewFromFloat(val1)
	fromFloat2 := decimal.NewFromFloat(val2)
	return fromFloat1.Cmp(fromFloat2)
}

// DecimalZero 0
func DecimalZero() float64 {
	v, _ := decimal.Zero.Float64()
	return v
}

// StrDecimalMultiplyRoundFixNum 相乘
func StrDecimalMultiplyRoundFixNum(val1 string, val2 float64, num int32) float64 {
	fromFloat1, _ := decimal.NewFromString(val1)
	fromFloat2 := decimal.NewFromFloat(val2)
	v, _ := fromFloat1.Mul(fromFloat2).Round(num).Float64()
	return v
}
