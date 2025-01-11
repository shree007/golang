package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	path := "dir1/dir2/filename.txt"
	pathJoin()
	getDirAndFile(path)
	splitPath(path)
	getExtention(path)
	detectAbsolutePath(path)
	walkIntoDirectory()
}

func walkIntoDirectory() {
	err := filepath.Walk("/Users/hero/src/golang", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func detectAbsolutePath(path string) {
	fmt.Println(filepath.IsAbs(path))
}

func getExtention(path string) {
	fmt.Println(filepath.Ext(path))
}

func splitPath(path string) {
	dir, file := filepath.Split(path)
	fmt.Println("Directory: ", dir)
	fmt.Println("File: ", file)
}

func pathJoin() {
	completePath := filepath.Join("dir1", "dir2", "filename.txt")
	fmt.Println(completePath)
}

func getDirAndFile(path string) {
	fmt.Println("except last element in path", filepath.Dir(path))
	fmt.Println("last element of path", filepath.Base(path))
}
