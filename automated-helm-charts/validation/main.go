package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ChartEntry struct {
	Metadata struct {
		Version string `yaml:"version"`
	} `yaml:"metadata"`
}

type indexFile struct {
	Entries map[string][]ChartEntry `yaml:"entries"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter index file path")
	indexFilePath, _ := reader.ReadString('\n')
	indexFilePath = strings.TrimSpace(indexFilePath)

	fmt.Println("Enter chart name")
	chartName, _ := reader.ReadString('\n')
	chartName = strings.TrimSpace(chartName)

	fmt.Println("Enter version of chart")
	chartVersion, _ := reader.ReadString('\n')
	chartVersion = strings.TrimSpace(chartVersion)

	matchedEntry, err := validateChartVersion(indexFilePath, chartName, chartVersion)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if matchedEntry != nil {
		log.Infof("selected version %s has been found", chartVersion)
	} else {
		log.Infof("selected version %s has NOT been found", chartVersion)
	}
}

func validateChartVersion(indexFilePath string, chartName string, chartVersion string) (*ChartEntry, error) {
	file, err := os.OpenFile(indexFilePath, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("Failed to open index file %v", err)
	}
	defer file.Close()

	var indexData indexFile
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&indexData); err != nil {
		return nil, fmt.Errorf("failed to parse index file: %w", err)
	}

	chartentries, ok := indexData.Entries[chartName]
	if !ok {
		log.Errorf("There is no chartName %s available in yaml files %v", chartName, err)
	}
	for _, entry := range chartentries {
		if entry.Metadata.Version == chartVersion {
			log.Infof("Version %s of ChartName %s is present in index file", chartVersion, chartName)
			return &entry, nil
		}
	}
	log.Infof("Version %s for chart %s not found.\n", chartVersion, chartName)
	return nil, nil
}
