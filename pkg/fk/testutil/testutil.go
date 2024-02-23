package testutil

import "fmt"

func Output(funcName string, value any) {
	fmt.Printf("%s: [%v]\n", funcName, value)
}
