package testutil

import (
	"fmt"
	"strconv"
	"strings"
)

func Output(funcName string, value any) {
	fmt.Printf("%s: [%v]\n", funcName, value)
}

func IsInSlice[T comparable](val T, slice []T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func GetDecimalLength(val float64) int {
	// FIXME: 末尾が0の場合に、切り捨てられるため、正確な桁数が取得できない
	strVal := strconv.FormatFloat(val, 'f', -1, 64)
	decimalVal := strings.Split(strVal, ".")
	if len(decimalVal) < 2 {
		return 0
	}
	return len(decimalVal[1])
}
