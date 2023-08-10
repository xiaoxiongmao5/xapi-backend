package utils

import (
	"errors"
	"strconv"
)

/** 字符串切片转成int切片
 */
func StrSlice2IntSlice(strSlice []string) ([]int, error) {
	if strSlice == nil {
		return nil, errors.New("param is nil")
	}
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		intVal, err := strconv.Atoi(str)
		if err != nil {
			return intSlice, err
		}
		intSlice[i] = intVal
	}
	return intSlice, nil
}

/** int切片转字符串切片
 */
func IntSlice2StrSlice(intSlice []int) []string {
	if intSlice == nil {
		return nil
	}
	strSlice := make([]string, len(intSlice))
	for i, intVal := range intSlice {
		strSlice[i] = strconv.Itoa(intVal)
	}
	return strSlice
}
