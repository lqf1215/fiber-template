package pkg

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
)

func CmpBigIntString(a, b string) int {
	aAmount, bAmount := convertToBigInt(a, b)
	result := aAmount.Cmp(&bAmount)
	switch result {
	case -1:
		return -1 //a小于b
	case 0:
		return 0 //a等于b
	case 1:
		return 1 //a大于b
	}
	return 2
}
func CmpBigFloatString(a, b string) int {
	aAmount, bAmount := convertToBigFloat(a, b)
	result := aAmount.Cmp(&bAmount)
	switch result {
	case -1:
		return -1 //a小于b
	case 0:
		return 0 //a等于b
	case 1:
		return 1 //a大于b
	}
	return 2
}
func BigIntStringSub(a, b string) string {
	aAmount, bAmount := convertToBigInt(a, b)
	return new(big.Int).Sub(&aAmount, &bAmount).String()
}
func BigIntStringAdd(a, b string) string {
	aAmount, bAmount := convertToBigInt(a, b)
	return new(big.Int).Add(&aAmount, &bAmount).String()
}
func BigIntStringMul(a, b string) string {
	aAmount, bAmount := convertToBigInt(a, b)
	return new(big.Int).Mul(&aAmount, &bAmount).String()
}
func BigIntStringQuo(a, b string) string {
	aAmount, bAmount := convertToBigInt(a, b)
	return new(big.Int).Quo(&aAmount, &bAmount).String()
}
func convertToBigInt(a, b string) (big.Int, big.Int) {
	var aAmount, bAmount big.Int
	aAmount.SetString(a, 10)
	bAmount.SetString(b, 10)
	return aAmount, bAmount
}
func convertToBigFloat(a, b string) (big.Float, big.Float) {
	var aAmount, bAmount big.Float
	aAmount.SetString(a)
	bAmount.SetString(b)
	return aAmount, bAmount
}
func convertStringToBigFloat(a string) big.Float {
	var aAmount big.Float
	aAmount.SetString(a)
	return aAmount
}
func BigFloat64StringQuo(a, b string) string {
	aAmountInEth, _ := decimal.NewFromString(a)
	bAmountInEth, _ := decimal.NewFromString(b)
	return aAmountInEth.Div(bAmountInEth).String()
}
func BigFloat64StringQuoTruncate4(a, b string) string {
	aAmountInEth, _ := decimal.NewFromString(a)
	bAmountInEth, _ := decimal.NewFromString(b)
	return aAmountInEth.Div(bAmountInEth).Truncate(4).String()
}
func BigFloat64StringQuoTruncateZero(a, b string) string {
	aAmountInEth, _ := decimal.NewFromString(a)
	bAmountInEth, _ := decimal.NewFromString(b)
	return aAmountInEth.Div(bAmountInEth).Truncate(0).String()
}

func BigFloat64StringMul(a, b string) string {
	aAmountInEth, _ := decimal.NewFromString(a)
	bAmountInEth, _ := decimal.NewFromString(b)
	return aAmountInEth.Mul(bAmountInEth).String()
}

func BigFloat64StringMulTruncateZero(a, b string) string {
	aAmountInEth, _ := decimal.NewFromString(a)
	bAmountInEth, _ := decimal.NewFromString(b)
	return aAmountInEth.Mul(bAmountInEth).Truncate(0).String()
}

func BigIntMulFloat64(a string, b float64) string {
	aAmountInEth, _ := decimal.NewFromString(a)
	bAmountInEth := decimal.NewFromFloat(b)
	return aAmountInEth.Mul(bAmountInEth).Truncate(0).String()
}
func ParseStringToBigInt(str string) *big.Int {
	// 将字符串表示的金额转换为大整数
	amountWei, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return nil
	}
	return amountWei
}
func ParasStringDecimal(a string) string {
	if strings.Contains(a, ".") {
		return strings.Split(a, ".")[0] + "." + strings.Split(a, ".")[1][:4]
	} else {
		floatValue, err := strconv.ParseFloat(a, 64)
		if err != nil {
			fmt.Printf("转换错误: %v\n", err)
			return a
		}
		normalStr := strconv.FormatFloat(floatValue, 'f', 4, 64)
		return normalStr
	}
}

func DecimalsToWei(amount *decimal.Decimal, decimals *big.Int) *big.Int {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromBigInt(decimals, 0))
	result := amount.Mul(mul)
	wei := new(big.Int)
	wei.SetString(result.String(), 10)
	return wei
}

func EthToWei(amount *decimal.Decimal) *big.Int {
	wei := DecimalsToWei(amount, big.NewInt(18))
	return wei
}

// SafeMulDivBigInt (amountBigInt, 40, new(big.Int).SetInt64(100)) //  fun(金额，40，100) = 金额*40%
func SafeMulDivBigInt(value, mulValue, divValue *big.Int) *big.Int {
	temp := new(big.Int)
	temp.Mul(value, mulValue)
	result := new(big.Int)
	result.Quo(temp, divValue)
	return result
}
