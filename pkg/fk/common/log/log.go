package log

import (
	"log"

	"github.com/yrichika/gfaker/pkg/fk/common"
)

func WrongUsage(msg string, skip int) {
	trueSkip := skip + 1
	caller, file, line := common.GetCallerInfo(trueSkip)
	log.Printf("gfaker might-be-wrong-usage warning. Empty value of the type returned: \"%s\" at [%s]: line [%d]: caller: [%s]:", msg, file, line, caller.Name())
}
