package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	pathJoin()
}

func pathJoin() {
	completePath := filepath.Join("dir1", "dir2", "filename.txt")
	fmt.Println(completePath)
}
