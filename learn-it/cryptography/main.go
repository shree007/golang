package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hashToString()
}

func hashToString() {
	input := "I am string"

	hash := sha256.Sum256([]byte(input))
	fmt.Printf("SHA-256 Hash: %x\n", hash)
}
