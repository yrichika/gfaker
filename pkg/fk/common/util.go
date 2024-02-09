package common

import (
	"runtime"
)

func GetCallerInfo(skip int) (string, int) {
	trueSkip := skip + 1
	_, file, line, _ := runtime.Caller(trueSkip)
	return file, line
}

// TODO: エラーメッセージを共通化させる関数
