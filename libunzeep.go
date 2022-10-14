package libunzeep

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"
)

func Open(filename string) ([]*zip.File, error) {
	var files []*zip.File
	unzipChannel, err := OpenChannel(filename)
	for file := range unzipChannel {
		files = append(files, file)
	}
	return files, err
}

func OpenChannel(filename string) (chan *zip.File, error) {
	unzipChannel := make(chan *zip.File)
	zipFiles, err := zip.OpenReader(filename)
	if err != nil {
		close(unzipChannel)
		return unzipChannel, CanNotReadZipError{err}
	}
	go func() {
		for _, zipFile := range zipFiles.File {
			for innerZipFile := range Unzeep(zipFile) {
				unzipChannel <- innerZipFile
			}
		}
		close(unzipChannel)
	}()
	return unzipChannel, err
}

func Unzeep(zipFile *zip.File) chan *zip.File {
	unzipChannel := make(chan *zip.File)
	go func() {
		if strings.HasSuffix(zipFile.Name, "zip") {
			buffer := new(bytes.Buffer)
			file, err := zipFile.Open()
			buffer.ReadFrom(file)
			file.Close()
			inMemoryFile := buffer.Bytes()
			inMemorySize := int64(len(inMemoryFile))
			inMemoryReader := bytes.NewReader(inMemoryFile)
			innerFiles, err := zip.NewReader(inMemoryReader, inMemorySize)
			if err != nil {
				unzipChannel <- zipFile
				close(unzipChannel)
				return
			}
			for _, innerFile := range innerFiles.File {
				innerFile.Name = fmt.Sprintf("%s/%s", zipFile.Name, innerFile.Name)
				innerZipFiles := Unzeep(innerFile)
				for innerZipFile := range innerZipFiles {
					unzipChannel <- innerZipFile
				}
			}
		} else {
			unzipChannel <- zipFile
		}
		close(unzipChannel)
	}()
	return unzipChannel
}
