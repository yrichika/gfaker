package testutil

import "fmt"

func Output(funcName string, value any) {
	fmt.Printf("%s: [%v]\n", funcName, value)
}

func IsInArray[T comparable](val T, array []T) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}
