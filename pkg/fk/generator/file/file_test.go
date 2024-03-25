package file

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
	"github.com/yrichika/gfaker/pkg/fk/provider/global"
)

func TestFile(testingT *testing.T) {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Files: global.CreateFiles(),
	}

	file := New(coreRand, global)

	t := gt.CreateTest(testingT)
	t.Describe("File", func() {
		t.Test("MimeType should return a string", func() {
			mimeType := file.MimeType()

			_, exists := global.Files.MimeTypes[mimeType]
			gt.Expect(t, &exists).ToBe(true)
		})

		t.Test("Extension should return a string", func() {
			extension := file.Extension()

			gt.Expect(t, &extension).ToMatchRegex(`^[0-9a-z-]{2,10}$`)
		})
	})

	t2 := gt.CreateTest(testingT)
	destDir := "../../../../tmp"
	t2.AfterAll(func() {
		// Clean up
		deleteFilesInDir(destDir)
	})
	t2.Describe("File", func() {
		content := "Hello, World!"
		extension := "txt"
		t2.Test("WriteWithText should create file and return a relative path if returnFullPath false", func() {
			returnFullPath := false
			filePath, err := file.WriteWithText(destDir, content, extension, returnFullPath)

			r := FileExists(filePath)
			gt.Expect(t2, &r).ToBe(true)
			gt.Expect(t2, &filePath).ToMatchRegex(`^..\/..\/..\/..\/tmp/[0-9a-z-]{36}\.txt$`)

			// This does not work because of type mismatch:
			// gt.Expect(t2, &err).ToBe(nil)
			errString := fmt.Sprint(err)
			gt.Expect(t2, &errString).ToBe("<nil>")
		})

		t2.Test("WriteWithText should return absolute path if returnFullPath true", func() {
			returnFullPath := true
			filePath, err := file.WriteWithText(destDir, content, extension, returnFullPath)

			r := FileExists(filePath)
			gt.Expect(t2, &r).ToBe(true)

			absPath, _ := filepath.Abs(filePath)
			gt.Expect(t2, &filePath).ToBe(absPath)

			errString := fmt.Sprint(err)
			gt.Expect(t2, &errString).ToBe("<nil>")
		})

		srcFilePath := "../../../../test_resources/files/sample.txt"
		t2.Test("CopyFrom should copy file from src to dest and should return relative path if returnFullPath is false", func() {
			returnFullPath := false
			extension := "txt"
			filePath, err := file.CopyFrom(destDir, srcFilePath, extension, returnFullPath)

			r := FileExists(filePath)
			gt.Expect(t2, &r).ToBe(true)
			gt.Expect(t2, &filePath).ToMatchRegex(`^..\/..\/..\/..\/tmp/[0-9a-z-]{36}\.txt$`)

			errString := fmt.Sprint(err)
			gt.Expect(t2, &errString).ToBe("<nil>")

		})

		t2.Test("CopyFrom should copy file from src to dest and should return fill path if returnFullPath is true", func() {
			returnFullPath := true
			extension := "txt"
			filePath, err := file.CopyFrom(destDir, srcFilePath, extension, returnFullPath)

			r := FileExists(filePath)
			gt.Expect(t2, &r).ToBe(true)

			absPath, _ := filepath.Abs(filePath)
			gt.Expect(t2, &filePath).ToBe(absPath)

			errString := fmt.Sprint(err)
			gt.Expect(t2, &errString).ToBe("<nil>")

		})
	})
}

func FileExists(filePath string) bool {
	_, pathErr := os.Stat(filePath)
	r := os.IsNotExist(pathErr) // DO NOT USE IsExist().
	return !r
}

func deleteFilesInDir(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
