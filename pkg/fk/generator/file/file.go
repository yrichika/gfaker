package file

import (
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type File struct {
	rand *core.Rand
	data *provider.Files
}

func New(rand *core.Rand, global *provider.Global) *File {
	return &File{
		rand,
		global.Files,
	}
}

func (f *File) MimeType() string {
	m, _ := f.rand.Map.KeySliceValue(f.data.MimeTypes)
	return m.(string)
}

func (f *File) Extension() string {
	_, e := f.rand.Map.KeySliceValue(f.data.MimeTypes)
	extensions := e.([]any)
	// compare to mime types slice, extension slices are usually very small.
	// it should not be any performance issue.
	strExtensions := make([]string, len(extensions))
	for i, v := range extensions {
		strExtensions[i] = v.(string)
	}
	return f.rand.Slice.StrElem(strExtensions)
}

// TODO:
func (f *File) File() string {
	return ""
}
