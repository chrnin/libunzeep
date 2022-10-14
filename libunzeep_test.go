package libunzeep

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Open_WithAbsentFile(t *testing.T) {
	ass := assert.New(t)
	filename := "nobodyThere.zip"
	expectedError := fmt.Sprintf("can not read zip file: open %s: no such file or directory", filename)

	// WHEN
	files, err := Open(filename)

	// THEN
	ass.EqualError(err, expectedError)
	count := 0
	for range files {
		count++
	}
	ass.Empty(count)
}

func Test_Open_WithNestedFile(t *testing.T) {
	ass := assert.New(t)
	filename := "test/test.zip"

	zipFiles, err := Open(filename)
	ass.Nil(err)
	ass.Len(zipFiles, 20)
}

func Test_Open_WithFakeZip(t *testing.T) {
	ass := assert.New(t)
	filename := "test/bad.zip"

	zipFiles, err := Open(filename)
	ass.ErrorAs(err, &CanNotReadZipError{})
	ass.Len(zipFiles, 0)
}

func Test_Open_WithFakeZipInside(t *testing.T) {
	ass := assert.New(t)
	filename := "test/zippedBad.zip"

	zipFiles, err := Open(filename)
	ass.Nil(err)
	ass.Len(zipFiles, 1)
}
