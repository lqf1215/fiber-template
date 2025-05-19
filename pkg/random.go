package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const alphanumerics = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alphanumericsUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRandomName(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// generates a pseudo-random string of length n.
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanumerics[rand.Int63()%int64(len(alphanumerics))]
	}
	Time := TimeNowUnixStr()
	bytes := bytes.Join([][]byte{[]byte(Time), b}, []byte{})
	hash := sha256.Sum256(bytes) // 对随机字节进行SHA256哈希处理
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}

func RandomUID(digits int) int64 {
	if digits < 1 {
		return 0 // 如果输入的位数小于 1，则返回 0
	}
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	// 计算最小值和最大值，确保 ID 始终具有指定的位数
	min := int(math.Pow10(digits - 1))
	max := int(math.Pow10(digits)) - 1

	// 生成并返回随机数
	return int64(rng.Intn(max-min+1) + min)
}

func RandomInt64() int64 {
	// 创建一个随机源
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	// 生成随机ID
	randomID := rng.Int63()
	return randomID
}

// 从30～120范围中随机返回一个值
func RandomRangeInt(min, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return rng.Intn(max-min) + min
}

func RandomRangeFloat(min, max float64) float64 {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return rng.Float64()*(max-min) + min
}

func RandomCodes(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// 生成随机大写字符串
func ReferralCodeGeneration(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanumericsUpper[rand.Int63()%int64(len(alphanumericsUpper))]
	}
	return string(b)
}

// RandomHashTradeNo 订单编号
func RandomHashTradeNo() string {
	return strings.ReplaceAll(time.Now().Format("060102150405.99999999"), ".", "")

}
