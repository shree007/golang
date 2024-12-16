package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "read-file.txt"
	fmt.Println("Reading file content as a string")
	readEntireFileAsStringContent(fileName)

	fmt.Println("Read file line by line")
	readFileLineByLine(fileName)

	fmt.Println("Read file by chunks")
	readFileByChunks(fileName)
}

func readFileByChunks(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file %v", err)
	}
	defer file.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Failed to read file: %v", err)
		}
		fmt.Println(string(buffer[:n]))
	}
}

func readFileLineByLine(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file %v", err)
	}
}

func readEntireFileAsStringContent(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	fmt.Println(string(content))
}
