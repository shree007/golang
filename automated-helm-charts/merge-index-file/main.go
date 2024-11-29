package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// https://zhwt.github.io/yaml-to-go/
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

func main() {
	var jfrogIndex HelmIndex
	err := loadYAML("indexfile-generated-by-helm-sdk.yaml", &jfrogIndex)
	if err != nil {
		log.Fatalf("Error loading yaml file %v", err)
	}
	fmt.Println(jfrogIndex)
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
