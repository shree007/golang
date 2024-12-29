package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type ExpectedIndexFile struct {
	Generated string                          `yaml:"generated"`
	Entries   map[string][]ExpectedChartEntry `yaml:"entries"`
}

type ExpectedChartEntry struct {
	ApiVersion  string       `yaml:"apiVersion"`
	Created     time.Time    `yaml:"created"`
	Description string       `yaml:"description"`
	Digest      string       `yaml:"digest"`
	Home        string       `yaml:"home"`
	Keywords    []string     `yaml:"keywords"`
	Maintainers []Maintainer `yaml:"maintainers"`
	Name        string       `yaml:"name"`
	Sources     []string     `yaml:"sources"`
	Urls        []string     `yaml:"urls"`
	AppVersion  string       `yaml:"appVersion"`
	Version     string       `yaml:"version"`
}

type Maintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
	URL   string `yaml:"url"`
}

type HelmIndex struct {
	Entries map[string][]ChartEntry `yaml:"entries"`
}

type ChartEntry struct {
	Metadata  ChartMetadata `yaml:"metadata"`
	URLs      []string      `yaml:"urls"`
	Created   time.Time     `yaml:"created"`
	Digest    string        `yaml:"digest"`
	Removed   bool          `yaml:"removed"`
	Checksum  string        `yaml:"checksumdeprecated"`
	TillerVer string        `yaml:"tillerversiondeprecated"`
}

type ChartMetadata struct {
	APIVersion  string       `yaml:"apiversion"`
	Name        string       `yaml:"name"`
	Home        string       `yaml:"home"`
	Sources     []string     `yaml:"sources"`
	Version     string       `yaml:"version"`
	Description string       `yaml:"description"`
	Keywords    []string     `yaml:"keywords"`
	Maintainers []Maintainer `yaml:"maintainers"`
	AppVersion  string       `yaml:"appversion"`
}

func main() {
	var jfrogIndex HelmIndex
	err := loadYAML("indexfile-generated-by-helm-sdk.yaml", &jfrogIndex)
	if err != nil {
		log.Fatalf("Error loading YAML file: %v", err)
	}

	expectedIndex := buildExpectedIndex(jfrogIndex)
	err = saveToYAML("final-expected-index.yaml", expectedIndex)
	if err != nil {
		log.Fatalf("Error saving to YAML: %v", err)
	}
}

func loadYAML(filePath string, data interface{}) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(fileContent, data)
}

func buildExpectedIndex(jfrogIndex HelmIndex) ExpectedIndexFile {
	expectedIndex := ExpectedIndexFile{
		Generated: time.Now().Format(time.RFC3339),
		Entries:   make(map[string][]ExpectedChartEntry),
	}

	for entryName, charts := range jfrogIndex.Entries {
		for _, chart := range charts {
			expectedChart := ExpectedChartEntry{
				ApiVersion:  chart.Metadata.APIVersion,
				Name:        chart.Metadata.Name,
				Created:     chart.Created,
				Description: chart.Metadata.Description,
				Digest:      chart.Digest,
				Home:        chart.Metadata.Home,
				Keywords:    chart.Metadata.Keywords,
				Maintainers: chart.Metadata.Maintainers,
				Sources:     chart.Metadata.Sources,
				Urls:        chart.URLs,
				AppVersion:  chart.Metadata.AppVersion,
				Version:     chart.Metadata.Version,
			}
			expectedIndex.Entries[entryName] = append(expectedIndex.Entries[entryName], expectedChart)
		}
	}
	return expectedIndex
}

func saveToYAML(filePath string, data interface{}) error {
	fileContent, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, fileContent, 0644)
}
