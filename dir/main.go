package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// Or using a variable:

func DirSize2(path string) (int64, error) {
	var size int64
	adjSize := func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	}
	err := filepath.Walk(path, adjSize)

	return size, err
}

func main() {
	dir_size,err := DirSize("../dir")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("dir目录大小:",dir_size)
	}

}