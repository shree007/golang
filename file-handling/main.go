package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	fileName := "read-file.txt"
	fmt.Println("Reading file content as a string")
	readEntireFileAsStringContent(fileName)

	fmt.Println("Read file line by line")
	readFileLineByLine(fileName)

	fmt.Println("Read file by chunks")
	readFileByChunks(fileName)
	log.Info("<<<<<<<<< CLOSE READING FILEs>>>>>>>>>")

	log.Info("<<<<<<<<<< OPEN WRITING FILEs>>>>>>>>>")

	writeIntoFileCompleteContent()
	writeIntoFileAppend()
	writeIntoFileAppendUsing()
	writeIntoCsv()

	log.Info("<<<<<<<<<< CLOSE WRITING FILEs>>>>>>>>>")

}

func writeIntoCsv() {
	file, err := os.Create("write-into-me.csv")
	if err != nil {
		log.Fatalf("Failed to create file %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	data := [][]string{
		{"Name", "Age", "Country"},
		{"Alice", "30", "USA"},
		{"Bob", "25", "Canada"},
		{"Charlie", "35", "UK"},
	}

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			log.Fatalf("Failed to write row to CSV: %v", err)
		}
	}

	log.Println("CSV file written successfully!")
}

func writeIntoFileAppendUsing() {
	fileName := "write-file-complete-content.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("Appending new content using bufio.Writer.\n")
	if err != nil {
		log.Fatalf("Failed to write buffered content: %v", err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatalf("Failed to flush buffer: %v", err)
	}

	log.Println("Buffered data appended successfully!")
}

func writeIntoFileAppend() {
	fileName := "write-file-complete-content.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	_, err = file.Write([]byte("\n Appending new content using Write.\n"))
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	log.Println("Data appended successfully")
}

func writeIntoFileCompleteContent() {
	content := []byte("Hello Blr, Hello Delhi, Hello Mumbai")
	fileName := "write-file-complete-content.txt"

	err := os.WriteFile(fileName, content, 0644)

	if err != nil {
		log.Fatalf("writing into file is failed %v", err)
	}

	fmt.Println("File has been written successfully")

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
