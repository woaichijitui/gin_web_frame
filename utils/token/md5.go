package token

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// GenerateSalt 为密码生成盐值
func GenerateSalt(length int) string {
	const alphanumeric = "abcdefghijklmnopqrstuvwxyz_0123456789_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, length)
	rand.Seed(int64(time.Now().UnixNano()))
	for i := range bytes {
		bytes[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}
	return string(bytes)
}

// Md5 计算文件hash值
func Md5(data []byte) string {
	// 创建一个新的MD5哈希对象
	hash := md5.New()

	// 写入数据到哈希对象
	hash.Write(data)

	// 计算哈希值
	hashInBytes := hash.Sum(nil)

	// 将字节数组转换为十六进制字符串
	hashInHex := hex.EncodeToString(hashInBytes)
	return hashInHex
}
