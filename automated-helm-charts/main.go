package main

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

const (
	chartBasePath = "helm/charts"
	packagePath   = "temp-helm-storage"
	indexFilePath = "temp-helm-storage/index.yaml"
)

/*
Cleanup index file old packaged charts and before starting
*/
func init() {
	matches, err := filepath.Glob(filepath.Join(packagePath, "*.tgz"))
	if err != nil {
		log.Fatalf("Error finding .tgz files: %v", err)
	}

	for _, match := range matches {
		if err := os.Remove(match); err != nil {
			log.Printf("Failed to remove file %s: %v", match, err)
		}
	}
	if err := os.Truncate(indexFilePath, 0); err != nil {
		log.Fatalf("Failed to truncate file: %v", err)
	}
}

func main() {
	log.Info("Starting Helm chart processing")

	index := createIndexFile(indexFilePath)

	entries, err := os.ReadDir(chartBasePath)
	if err != nil {
		log.Fatal("Error Reading chart directory: ", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			processChart(entry.Name(), index)
		} else {
			log.Infof("%s is not a directory", entry.Name())
		}
	}

	if err := writeIndexFile(index, indexFilePath); err != nil {
		log.Fatal("Error writing index file: ", err)
	}
}

func processChart(chartName string, index *repo.IndexFile) {
	chartPath := filepath.Join(chartBasePath, chartName)
	log.Infof("Loading chart from %s", chartPath)

	chart, err := loadingChart(chartPath)
	if err != nil {
		log.Errorf("Error loading chart: %v", err)
		return
	}

	if err := updateDependencies(chartPath); err != nil {
		log.Errorf("Error updating dependencies: %v", err)
		return
	}

	if err := lintChart(chart); err != nil {
		log.Errorf("Linting Errors found in chart %s: %v", chart.Name(), err)
		return
	}

	chartURL := packageChart(chart)
	addToIndex(chart, chartURL, index)
}

func loadingChart(charPath string) (*chart.Chart, error) {
	return loader.Load(charPath)
}

func createIndexFile(indexFilePath string) *repo.IndexFile {
	index := repo.NewIndexFile()

	if _, err := os.Stat(indexFilePath); err == nil {
		indexFile, err := os.ReadFile(indexFilePath)
		if err != nil {
			log.Fatalf("Error reading index file: %v", err)
		}
		if err := yaml.Unmarshal(indexFile, index); err != nil {
			log.Fatalf("Error unmarshaling index file: %v", err)
		}
		log.Infof("Loaded existing index file from %s", indexFilePath)
	} else {
		log.Infof("No existing index file found. A new one will be created at %s", indexFilePath)
	}
	return index
}

func updateDependencies(chartPath string) error {
	settings := cli.New()
	manager := &downloader.Manager{
		ChartPath:  chartPath,
		Getters:    getter.All(settings),
		SkipUpdate: false,
		Out:        os.Stdout,
	}
	if err := manager.Update(); err != nil {
		return err
	}
	log.Info("Dependencies updated successfully")
	return manager.Build()
}

func BuildDependency(chartPath string) error {
	settings := cli.New()
	manager := &downloader.Manager{
		ChartPath:  chartPath,
		Getters:    getter.All(settings),
		SkipUpdate: false,
		Out:        os.Stdout,
	}
	return manager.Build()
}

func lintChart(chart *chart.Chart) error {
	if chart.Metadata == nil {
		return fmt.Errorf("missing chart metadata")
	}

	if chart.Metadata.Name == "" {
		return fmt.Errorf("chart name is missing")
	}

	if chart.Metadata.Version == "" {
		return fmt.Errorf("chart version is missing")
	}

	if len(chart.Templates) == 0 {
		return fmt.Errorf("chart %s has no templates", chart.Name())
	}

	if len(chart.Values) == 0 {
		log.Warnf("Chart %s has no default values", chart.Name())
	}

	log.Infof("Chart %s passed basic lint checks", chart.Name())
	return nil
}

func addToIndex(chart *chart.Chart, chartURL string, index *repo.IndexFile) {
	if existingVersions, ok := index.Entries[chart.Metadata.Name]; ok {
		for _, v := range existingVersions {
			if v.Version == chart.Metadata.Version {
				log.Infof("Chart %s version %s already exists in the index, skipping", chart.Metadata.Name, chart.Metadata.Version)
				return
			}
		}
	}
	index.MustAdd(chart.Metadata, chartURL, " ", " ")
	log.Infof("Added chart %s version %s to index", chart.Metadata.Name, chart.Metadata.Version)
}

func writeIndexFile(index *repo.IndexFile, path string) error {
	return index.WriteFile(path, 0644)
}

func packageChart(chart *chart.Chart) string {
	pkgPath, err := chartutil.Save(chart, packagePath)
	if err != nil {
		log.Errorf("error saving packaged chart %s: %v", chart.Name(), err)
		return ""
	}
	log.Infof("Packaged chart %s to %s", chart.Name(), pkgPath)
	return fmt.Sprintf("%s/%s-%s.tgz", packagePath, chart.Metadata.Name, chart.Metadata.Version)
}
