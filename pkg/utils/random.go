package utils

import (
	"math/rand"
	"strings"
	"time"
)

const nums = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomTxnID(n int) string {
	var sb strings.Builder
	k := len(nums)

	for i := 0; i < n; i++ {
		c := nums[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
