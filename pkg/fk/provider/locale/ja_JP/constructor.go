package ja_JP

import (
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/global"
)

func New() *base.Localized {
	return &base.Localized{
		People: CreatePeople(),
		Colors: global.CreateColors(),
	}
}
