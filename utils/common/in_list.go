package common

import "strings"

// InList 判断一个字符是否存在于字符切片里
func InList(list []string, str string) bool {
	for _, lstr := range list {
		if strings.ToLower(lstr) == str {
			return true
		}
	}
	return false
}
