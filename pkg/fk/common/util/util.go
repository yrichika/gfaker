package util

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"text/template"
	"time"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
)

// []string, []intなどの配列を[]anyに変換する
// 便利だが、パフォーマンスコストがかかるので、使いどころに注意
func ConvertToAnySlice[S ~[]E, E any](s S) []any {
	r := make([]any, len(s))
	for i, e := range s {
		r[i] = e
	}
	return r
}

// フォーマットの文字列をテンプレートとして扱い、データを埋め込み、文字列を返す
func RenderTemplate(format string, data interface{}) string {
	errMsgTmpl := "Failed to render locale template: %v"
	tmpl, err := template.New(format).Parse(format)
	if err != nil {
		errMsg := fmt.Sprintf(errMsgTmpl, err)
		log.GeneralError(errMsg, 2)
		return ""
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		errMsg := fmt.Sprintf(errMsgTmpl, err)
		log.GeneralError(errMsg, 2)
		return ""
	}

	return buf.String()
}

func RandSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TruncateToPrecision(val float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Trunc(val*multiplier) / multiplier
}

func CapFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
