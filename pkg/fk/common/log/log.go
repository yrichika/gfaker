package log

import (
	"fmt"
	"log"
	"runtime"
)

func GetCallerInfo(skip int) (*runtime.Func, string, int) {
	trueSkip := skip + 1
	_, file, line, _ := runtime.Caller(trueSkip)
	pc, _, _, _ := runtime.Caller(skip) // the function/method itself
	called := runtime.FuncForPC(pc)
	return called, file, line
}

func UnavailableLocaleMsg() string {
	return "This method is not available for the current locale."
}

func UnavailableLocale(skip int) {
	WrongUsage(UnavailableLocaleMsg(), skip+1)
}

func WrongUsageMsg(msg string) string {
	return fmt.Sprintf("gfaker might-be-wrong-usage warning. Empty value of the type returned: \"%s\"", msg)
}

func WrongUsage(msg string, skip int) {
	GeneralError(WrongUsageMsg(msg), skip+1)
}

func GeneralError(msg string, skip int) {
	trueSkip := skip + 1
	caller, file, line := GetCallerInfo(trueSkip)
	log.Printf("%s at [%s]: line [%d]: called: [%s]:", msg, file, line, caller.Name())
}
