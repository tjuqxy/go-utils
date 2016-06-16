package utils

import (
	"time"
	"math/rand"
)

var (
	r *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 包括传入的最大值, int值最小为1
func RandInt(maxInt int) int {
	return r.Intn(maxInt) + 1
}

func RandFloat(maxFloat int) float64 {
	return float64(maxFloat) * r.Float64()
}


func RandStringWithCertainLen(length int) string {
	buff := make([]byte, 0)
	for {
		if length <= 0 {
			break
		}
		buff = append(buff, byte(r.Intn(26) + 'a'))
		length--
	}
	return string(buff)
}

func RandString(maxLen int) string {
	length := RandInt(maxLen) + 10
    return RandStringWithCertainLen(length)
}

func RandRuneString(maxLen int) string {
	buff := make([]rune, 0)
	length := RandInt(maxLen) + 10
	for {
		if length < 0 {
			break
		}
		buff = append(buff, rune(r.Intn(0xFFFF)))
		length--
	}
	return string(buff)
}

func AllRuneSetString() string {
	ind  := 0
	buff := make([]rune, 65536)
	for ind < 65536 {
		buff[ind] = rune(ind)
		ind++
	}
	return string(buff)
}
