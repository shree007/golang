package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	pathJoin()
	getDirAndFile()
	splitPath()
}

func splitPath() {
	path := "dir1/dir2/filename.txt"
	dir, file := filepath.Split(path)
	fmt.Println("Directory: ", dir)
	fmt.Println("File: ", file)
}

func pathJoin() {
	completePath := filepath.Join("dir1", "dir2", "filename.txt")
	fmt.Println(completePath)
}

func getDirAndFile() {
	p := "dir1/dir2/filename.txt"
	fmt.Println("except last element in path", filepath.Dir(p))
	fmt.Println("last element of path", filepath.Base(p))
}
