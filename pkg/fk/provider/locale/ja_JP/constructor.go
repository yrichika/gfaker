package ja_JP

import (
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

func New() *provider.Localized {
	return &provider.Localized{
		People:    CreatePeople(),
		Addresses: CreateAddresses(),
		Companies: CreateCompanies(),
	}
}
