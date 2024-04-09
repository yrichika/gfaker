package lorem

import (
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Lorem struct {
	rand *core.Rand
	data *provider.Lorems
}

func New(rand *core.Rand, local *provider.Localized) *Lorem {
	return &Lorem{
		rand,
		local.Lorems,
	}
}

// WORKING:
