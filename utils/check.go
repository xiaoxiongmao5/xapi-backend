package utils

import (
	"reflect"
	"strings"
	glog "xj/xapi-backend/g_log"
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

// 检查是否一样（使用 == 检查）
func CheckSame[T string | int](desc string, str1 T, str2 T) bool {
	glog.Log.Infof("======= %s =======", desc)
	if str1 == str2 {
		glog.Log.Infof("相同 export: %v got: %v", str2, str1)
		return true
	} else {
		glog.Log.Errorf("不同 export: %v got: %v", str2, str1)
		return false
	}
}

// 检查字符串忽略大小写后是否一样（使用 EqualFold 检查）
func CheckSameStrFold(desc string, str1 string, str2 string) bool {
	glog.Log.Infof("======= %s =======", desc)
	if strings.EqualFold(str1, str2) {
		glog.Log.Infof("相同(已忽略大小写) export: %v got: %v", str2, str1)
		return true
	} else {
		glog.Log.Errorf("不同(已忽略大小写) export: %v got: %v", str2, str1)
		return false
	}
}

// 检查数组是否一样（使用 DeepEqual 检查）
func CheckSameArr[T string | int | []int](desc string, str1 T, str2 T) bool {
	glog.Log.Infof("======= %s =======", desc)
	if reflect.DeepEqual(str1, str2) {
		glog.Log.Infof("相同 export: %v got: %v", str2, str1)
		return true
	} else {
		glog.Log.Errorf("不同 export: %v got: %v", str2, str1)
		return false
	}
}
