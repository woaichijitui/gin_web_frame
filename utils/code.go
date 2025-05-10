package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Code() string {
	maxInt := big.NewInt(10000) // 4位数最大值10000（不包括10000）
	n, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%04d", n.Int64())
}
