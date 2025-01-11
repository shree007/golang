package main

import (
	"fmt"
	"io/fs"
	"os"
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
	cleanThePath()
	createDir()
	createNestedDir()
	removeDir()
	removeNestedDir()
}

func removeNestedDir() {
	err := os.RemoveAll("parent")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func removeDir() {
	err := os.Remove("newdir")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func createNestedDir() {
	err := os.MkdirAll("parent/child", 0755)
	if err != nil {
		fmt.Println(err)
	}
}

func createDir() {
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func cleanThePath() {
	path := "dir1/dir2/../../filename.txt"
	fmt.Println("original path: ", path)
	fmt.Println("Clean up ", filepath.Clean(path))
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
