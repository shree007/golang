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
	log.Info("Read Base Directory: ", chartBasePath)
	entries, err := os.ReadDir(chartBasePath)

	if err != nil {
		log.Error("Error Reading chart directory", err)
		os.Exit(1)
	}

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

	for _, entry := range entries {
		if entry.IsDir() {
			chartPath := filepath.Join(chartBasePath, entry.Name())
			log.Infof("Loading chart from %s", chartPath)
			chart, err := loadingChart(chartPath)

			if err != nil {
				log.Error("Erros in loading charts: ", err)
				continue
			}
			log.Infof("Loaded charts are %s", chart.Name())

			if err := updateDependency(chartPath); err != nil {
				log.Errorf("Error while updating dependencies: %v", err)
			}
			log.Infof("Dependencies built successfully for chart %s", chart.Name())

			if err := lintChart(chart); err != nil {
				log.Errorf("Linting Errors found in chart %s: %v", chart.Name(), err)
			} else {
				log.Infof("Linting passed for chart %s", chart.Name())
			}

			pkgPath, err := chartutil.Save(chart, packagePath)
			if err != nil {
				log.Errorf("error saving packaged chart %s: %v", chart.Name(), err)
			}
			log.Infof("Packaged chart %s to %s", chart.Name(), pkgPath)

			chartURL := fmt.Sprintf("%s/%s-%s.tgz", packagePath, chart.Metadata.Name, chart.Metadata.Version)
			log.Info("chartURL: ", chartURL)
			if existingVersions, ok := index.Entries[chart.Metadata.Name]; ok {
				versionExists := false
				for _, v := range existingVersions {
					if v.Version == chart.Metadata.Version {
						versionExists = true
						break
					}
				}
				if versionExists {
					log.Infof("Chart %s version %s already exists in the index, skipping", chart.Metadata.Name, chart.Metadata.Version)
					continue
				}
			}
			index.MustAdd(chart.Metadata, chartURL, " ", " ")
			log.Infof("Added chart %s version %s to index", chart.Metadata.Name, chart.Metadata.Version)

		} else {
			log.Infof("%s is not a directory: ", entry.Name())
		}
	}
	if err := index.WriteFile(indexFilePath, 0644); err != nil {
		log.Errorf("Error writing index file: %v", err)
	} else {
		log.Infof("Index file written successfully at %s", indexFilePath)
	}
}

func loadingChart(charPath string) (*chart.Chart, error) {
	return loader.Load(charPath)
}

func updateDependency(chartPath string) error {
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
