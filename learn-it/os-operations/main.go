package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	pathJoin()
	isDir()
}

func pathJoin() {
	completePath := filepath.Join("dir1", "dir2", "filename.txt")
	fmt.Println(completePath)
}

func isDir() {
	p := "dir1/dir2/filename.txt"
	fmt.Println("except last element in path", filepath.Dir(p))
	fmt.Println("last element of path", filepath.Base(p))
}
