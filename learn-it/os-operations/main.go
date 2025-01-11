package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "dir1/dir2/filename.txt"
	pathJoin()
	getDirAndFile(path)
	splitPath(path)
	getExtention(path)
	detectAbsolutePath(path)
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
