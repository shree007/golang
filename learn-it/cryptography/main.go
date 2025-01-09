package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	hashToString()
	hashFile()
	comparingHash()
}

func comparingHash() {
	data1 := "hello"
	data2 := "world"

	hash1 := sha256.Sum256([]byte(data1))
	hash2 := sha256.Sum256([]byte(data2))

	if hash1 == hash2 {
		fmt.Println("Hashes matched")
	} else {
		fmt.Println("Hashes unmatched")
	}
}

func hashFile() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		fmt.Println("Error hashing file:", err)
		return
	}
	hashSum := hasher.Sum(nil)
	fmt.Printf("File SHA-256 Hash: %x\n", hashSum)

}

func hashToString() {
	input := "I am string"

	hash := sha256.Sum256([]byte(input))
	fmt.Printf("SHA-256 Hash: %x\n", hash)
}
