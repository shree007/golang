package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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
	var finalIndex ExpectedIndexFile
	err := loadYAML("final-expected-index.yaml", &finalIndex)
	if err != nil {
		log.Printf("final-expected-index.yaml not found or failed to load: %v", err)
		finalIndex = ExpectedIndexFile{
			Generated: time.Now().Format(time.RFC3339),
			Entries:   make(map[string][]ExpectedChartEntry),
		}
	}

	var jfrogIndex HelmIndex
	err = loadYAML("indexfile-generated-by-helm-sdk.yaml", &jfrogIndex)
	if err != nil {
		log.Fatalf("Error loading YAML file: %v", err)
	}

	mergeIndexes(&finalIndex, jfrogIndex)

	err = saveToYAML("final-expected-index.yaml", finalIndex)
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

func saveToYAML(filePath string, data interface{}) error {
	fileContent, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, fileContent, 0644)
}

func mergeIndexes(finalIndex *ExpectedIndexFile, jfrogIndex HelmIndex) {
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

			existingEntries := finalIndex.Entries[entryName]
			isVersionExists := false

			for _, existingChart := range existingEntries {
				if existingChart.Version == expectedChart.Version {
					isVersionExists = true
					break
				}
			}

			if !isVersionExists {
				finalIndex.Entries[entryName] = append(finalIndex.Entries[entryName], expectedChart)
			}

			sort.SliceStable(finalIndex.Entries[entryName], func(i, j int) bool {
				return compareVersions(finalIndex.Entries[entryName][i].Version, finalIndex.Entries[entryName][j].Version) < 0
			})
		}
	}

	finalIndex.Generated = time.Now().Format(time.RFC3339)
}

func compareVersions(version1, version2 string) int {
	v1Parts := strings.Split(version1, ".")
	v2Parts := strings.Split(version2, ".")

	for i := 0; i < len(v1Parts) || i < len(v2Parts); i++ {
		var v1Part, v2Part int
		if i < len(v1Parts) {
			v1Part, _ = strconv.Atoi(v1Parts[i])
		}
		if i < len(v2Parts) {
			v2Part, _ = strconv.Atoi(v2Parts[i])
		}

		if v1Part < v2Part {
			return -1
		} else if v1Part > v2Part {
			return 1
		}
	}
	return 0
}
