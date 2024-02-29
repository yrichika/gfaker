package testutil

import "fmt"

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
