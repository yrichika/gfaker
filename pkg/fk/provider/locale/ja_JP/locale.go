package ja_JP

import (
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
)

func New() *base.Localized {
	return &base.Localized{
		FirstNameMale:       firstNameMale,
		FirstNameFemale:     firstNameFemale,
		LastName:            lastName,
		FirstKanaNameMale:   firstKanaNameMale,
		FirstKanaNameFemale: firstKanaNameFemale,
		LastKanaName:        lastKanaName,
	}
}
