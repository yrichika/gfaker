package file

import (
	"io"
	"os"
	"path/filepath"

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

// Create a file with the given content and extension.
// File name will be a UUID.
// extension should be without "."
func (f *File) WriteWithText(
	destDir string,
	content string,
	extension string,
	returnFullPath bool,
) (string, error) {

	DirErr := os.MkdirAll(destDir, 0777)
	if DirErr != nil {
		return "", DirErr
	}

	// 16 letters
	fileName := f.rand.Str.AlphaDigitsLike("****************") + "." + extension
	filePath := filepath.Join(destDir, fileName)

	data := []byte(content)
	FileErr := os.WriteFile(filePath, data, 0777)
	if FileErr != nil {
		return "", FileErr
	}

	if returnFullPath {
		return filepath.Abs(filePath)
	}

	return filePath, nil
}

// Create a file and copy the content from the source file.
// File name will be a UUID.
// Specify `srcFilePath` the path to the source file. Not directory name.
// extension should be without "."
func (f *File) CopyFrom(
	destDir string,
	srcFilePath string,
	extension string,
	returnFullPath bool,
) (string, error) {
	srcFile, openErr := os.Open(srcFilePath)
	if openErr != nil {
		return "", openErr
	}
	defer srcFile.Close()

	DirErr := os.MkdirAll(destDir, 0777)
	if DirErr != nil {
		return "", DirErr
	}
	// 16 letters
	fileName := f.rand.Str.AlphaDigitsLike("****************") + "." + extension
	filePath := filepath.Join(destDir, fileName)
	destFile, createErr := os.Create(filePath)
	if createErr != nil {
		return "", createErr
	}
	defer destFile.Close()

	_, copyErr := io.Copy(destFile, srcFile)
	if copyErr != nil {
		return "", copyErr
	}

	if returnFullPath {
		return filepath.Abs(filePath)
	}

	return filePath, nil
}
