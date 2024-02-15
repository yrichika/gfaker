package common

import (
	"bytes"
	"runtime"
	"text/template"
)

func GetCallerInfo(skip int) (*runtime.Func, string, int) {
	trueSkip := skip + 1
	pc, file, line, _ := runtime.Caller(trueSkip)
	caller := runtime.FuncForPC(pc)
	return caller, file, line
}

// []string, []intなどの配列を[]anyに変換する
// 便利だが、パフォーマンスコストがかかるので、使いどころに注意
func ConvertToAnyArray[S ~[]E, E any](s S) []any {
	r := make([]any, len(s))
	for i, e := range s {
		r[i] = e
	}
	return r
}

// フォーマットの文字列をテンプレートとして扱い、データを埋め込み、文字列を返す
func RenderTemplate(format string, data interface{}) (string, error) {
	tmpl, err := template.New("name").Parse(format)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
