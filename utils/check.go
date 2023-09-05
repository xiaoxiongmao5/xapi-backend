package utils

import (
	"fmt"
	"reflect"
)

// 检查是否为空字符串
func AreEmptyStrings(values ...string) bool {
	for _, value := range values {
		if value == "" {
			return true
		}
	}
	return false
}

// 检查是否一样（使用 DeepEqual 检查）
func CheckSame[T string | int | []int](desc string, str1 T, str2 T) bool {
	fmt.Printf("======= %s =======\n", desc)
	if reflect.DeepEqual(str1, str2) {
		fmt.Printf("\n相同\n export\t%v\n got\t%v\n", str2, str1)
		return true
	} else {
		fmt.Printf("\n不同\n export\t%v\n got\t%v\n", str2, str1)
		return false
	}
}
