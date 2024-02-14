package ja_JP

import (
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
)

func New() *base.Localized {
	return &base.Localized{
		Person: CreatePerson(),
	}
}
