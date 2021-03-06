package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// 切片去重
func SliceDuplicateRemoval(elements []string) (result []string) {
	tempMap := make(map[string]string)
	for _, v := range elements {
		tempMap[v] = v
	}
	for _, v := range tempMap {
		result = append(result, v)
	}
	return
}

func GetSafetySqlOr(fields []string) string {
	var safetyFields []string
	for _, f := range fields {
		safetyFields = append(safetyFields, fmt.Sprintf("code = '%s'", f))
	}
	return fmt.Sprintf("( %s )", strings.Join(safetyFields, " OR "))
}


func IsContainStr(li []string, x string) bool {
	for _, v := range li{
		if v == x{
			return true
		}
	}
	return false
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}