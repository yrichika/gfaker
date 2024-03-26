package global

import "github.com/yrichika/gfaker/pkg/fk/provider"

func New() *provider.Global {
	return &provider.Global{
		Colors: CreateColors(),
		Files:  CreateFiles(),
		Images: CreateImages(),
	}
}
