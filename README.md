# libunzeep
go library to survive with nested zip files

## get zip files in an array
```go
package libunzeep

import (
	"fmt"
	"github.com/chrnin/libunzeep"
)

func main() {
	zipFiles, err := libunzeep.Open("foo/bar.zip")
	if err != nil {
		panic(err)
	}
	for _, file := range zipFiles {
		fmt.Println(file.Name)
		reader, err := file.Open()
		if err != nil {
			panic(err)
		}
		fmt.Println(reader)
	}
}
```

## get zip files in a channel with a filename (memory friendly)
```go
package main 

import (
    "fmt"
    "github.com/chrnin/libunzeep"
)

func main() {
	zipFilesChannel, err := libunzeep.OpenChannel("foo/bar.zip")
	if err != nil {
		panic(err)
	}
	for file := range zipFilesChannel {
		fmt.Println(file.Name)
		reader, err := file.Open()
		if err != nil {
			panic(err)
		}
		fmt.Println(reader)
	}
}
```
