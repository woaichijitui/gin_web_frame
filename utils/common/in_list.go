package common

import (
	"cmp"
	"strings"
)

// InList 判断一个字符是否存在于字符切片里
func InList(list []string, str string) bool {
	for _, lstr := range list {
		if strings.ToLower(lstr) == str {
			return true
		}
	}
	return false
}

// 切片去重升级版 泛型参数 利用map的key不能重复的特性+append函数  一次for循环搞定
func ListUnique[T cmp.Ordered](ss []T) []T {
	size := len(ss)
	if size == 0 {
		return []T{}
	}
	newSlices := make([]T, 0) //这里新建一个切片,大于为0, 因为我们不知道有几个非重复数据,后面都使用append来动态增加并扩容
	m1 := make(map[T]byte)
	for _, v := range ss {
		if _, ok := m1[v]; !ok { //如果数据不在map中,放入
			m1[v] = 1                        // 保存到map中,用于下次判断
			newSlices = append(newSlices, v) // 将数据放入新的切片中
		}
	}
	return newSlices
}

// CompareSlices 对比两个切片，返回重复的元素切片和第一个切片中不存在但第二个切片存在的元素切片
func CompareSlices[T comparable](slice1, slice2 []T) (commonElements []T, diffElements []T) {
	// 用于记录第一个切片中的元素
	exists := make(map[T]bool)
	for _, item := range slice1 {
		exists[item] = true
	}

	for _, item := range slice2 {
		if exists[item] {
			// 如果元素存在于第一个切片中，添加到重复元素切片
			commonElements = append(commonElements, item)
		} else {
			// 如果元素不存在于第一个切片中，添加到差异元素切片
			diffElements = append(diffElements, item)
		}
	}
	return
}
