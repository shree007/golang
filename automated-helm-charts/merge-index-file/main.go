package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// https://zhwt.github.io/yaml-to-go/

// Struct for Input helm sdk generated
type HelmIndex struct {
	ServerInfo  interface{}             `yaml:"serverinfo"`
	APIVersion  string                  `yaml:"apiversion"`
	Generated   time.Time               `yaml:"generated"`
	Entries     map[string][]ChartEntry `yaml:"entries"`
	PublicKeys  []interface{}           `yaml:"publickeys"`
	Annotations map[string]interface{}  `yaml:"annotations"`
}

type ChartEntry struct {
	Metadata               ChartMetadata `yaml:"metadata"`
	URLs                   []string      `yaml:"urls"`
	Created                time.Time     `yaml:"created"`
	Removed                bool          `yaml:"removed"`
	Digest                 string        `yaml:"digest"`
	ChecksumDeprecate      string        `yaml:"checksumdeprecated"`
	EngineDeprecate        string        `yaml:"enginedeprecated"`
	TillerVersionDeprecate string        `yaml:"tillerversiondeprecated"`
	URLDeprecate           string        `yaml:"urldeprecated"`
}

type ChartMetadata struct {
	Name         string                 `yaml:"name"`
	Home         string                 `yaml:"home"`
	Sources      []string               `yaml:"sources"`
	Version      string                 `yaml:"version"`
	Description  string                 `yaml:"description"`
	Keywords     []string               `yaml:"keywords"`
	Maintainers  []Maintainer           `yaml:"maintainers"`
	Icon         string                 `yaml:"icon"`
	APIVersion   string                 `yaml:"apiversion"`
	Condition    string                 `yaml:"condition"`
	Tags         string                 `yaml:"tags"`
	AppVersion   string                 `yaml:"appversion"`
	Deprecated   bool                   `yaml:"deprecated"`
	Annotations  map[string]interface{} `yaml:"annotations"`
	KubeVersion  string                 `yaml:"kubeversion"`
	Dependencies []string               `yaml:"dependencies"`
	Type         string                 `yaml:"type"`
}

type Maintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
	URL   string `yaml:"url"`
}

// Expected struct
type ExpectedIndexFile struct {
	Generated string                          `yaml:"generated"`
	Entries   map[string][]ExpectedChartEntry `yaml:"entries"`
}

type ExpectedChartEntry struct {
	ApiVersion  string               `yaml:"apiVersion"`
	Created     string               `yaml:"created"`
	Description string               `yaml:"description"`
	Digest      string               `yaml:"digest"`
	Home        string               `yaml:"home"`
	Keywords    []string             `yaml:"keywords"`
	Maintainers []ExpectedMaintainer `yaml:"maintainers"`
	Name        string               `yaml:"name"`
	Sources     []string             `yaml:"sources"`
	Urls        []string             `yaml:"urls"`
	AppVersion  string               `yaml:"appVersion"`
	Version     string               `yaml:"version"`
}

type ExpectedMaintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func main() {
	var jfrogIndex HelmIndex
	err := loadYAML("indexfile-generated-by-helm-sdk.yaml", &jfrogIndex)
	if err != nil {
		log.Fatalf("Error loading yaml file %v", err)
	}
	buildExpectedchartEntry(jfrogIndex)
}

func loadYAML(filePath string, data interface{}) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileContent, data)
	if err != nil {
		return err
	}
	return nil
}

// appVersion: v0.12.0
// version: 4.11.0
func buildExpectedchartEntry(jfrogIndex HelmIndex) {
	// First I can print all those needed data
	for entryName, Charts := range jfrogIndex.Entries {
		fmt.Println("Chart Entry Name", entryName)
		for _, chart := range Charts {
			fmt.Println("apiVersion: ", chart.Metadata.APIVersion)
			fmt.Println("created: ", chart.Created)
			fmt.Println("description: ", chart.Metadata.Description)
			fmt.Println("digest: ", chart.Digest)
			fmt.Println("home: ", chart.Metadata.Home)
			fmt.Println("keywords", chart.Metadata.Keywords)
			fmt.Println("maintainers", chart.Metadata.Maintainers)
			fmt.Println("sources", chart.Metadata.Sources)
			fmt.Println("urls", chart.URLs)
			fmt.Println("appVersion", chart.Metadata.AppVersion)
			fmt.Println("version", chart.Metadata.Version)
		}
	}

}
