package randpwd

import (
	"strings"
)

func GetEasyPassword(num int) string {
	var result []string
	for k, _ := range Easy {
		result = append(result, k)
	}
	return strings.Join(result[:num], "")
}

func GetHardPassword(num int) string {
	var result []string
	for k, _ := range Data {
		result = append(result, k)
	}
	return strings.Join(result[:num], "")
}
