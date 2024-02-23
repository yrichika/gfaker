package en_US

import (
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

func New() *provider.Localized {
	return &provider.Localized{
		People: CreatePeople(),
	}
}
