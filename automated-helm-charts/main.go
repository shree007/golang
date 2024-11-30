package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

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
	chartBasePath    = "helm/charts"
	packageOutputDir = "temp-helm-storage"
	jfrogArtifactURL = "https://linkinpark.jfrog.io/artifactory/"
	jfrogRepoName    = "linkinpark-helmchart-helm-local"
	indexFilePath    = "temp-helm-storage/index.yaml"
)

type chartDetails struct {
	APIVersion  string              `yaml:"apiVersion"`
	Created     string              `yaml:"created"`
	Description string              `yaml:"description"`
	Digest      string              `yaml:"digest"`
	Home        string              `yaml:"home"`
	Keywords    []string            `yaml:"keywords,omitempty"`
	Maintainers []*chart.Maintainer `yaml:"maintainers, omitempty"`
	Name        string              `yaml:"name"`
	Sources     []string            `yaml:"sources,omitempty"`
	URLs        []string            `yaml:"urls"`
	AppVersion  string              `yaml:"appVersion"`
	Version     string              `yaml:"version"`
}

var packagedChartPaths []string

func init() {
	if _, err := os.Stat(packageOutputDir); os.IsNotExist(err) {
		log.Errorf("Directory does not exists, nothing to remove %v", err)
	}

	if err := os.RemoveAll(packageOutputDir); err != nil {
		log.Errorf("Failed to remove directory %v", err)
	} else {
		log.Info("Directory has been removed")
	}
}

func main() {
	log.Info("Processing Helm chart...")

	// if err := downloadIndexFile(); err != nil {
	// 	log.Fatalf("Failed to Download Index file %v", err)
	// }

	entries, err := os.ReadDir(chartBasePath)
	if err != nil {
		log.Fatal("Error Reading chart directory: ", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			processChart(entry.Name())
		} else {
			log.Infof("%s is not a directory", entry.Name())
		}
	}

	index, err := loadOrCreateIndex(indexFilePath)
	if err != nil {
		log.Fatalf("Failed to load or create index file: %v", err)
	}
	fmt.Println(index)

	fmt.Println(packagedChartPaths)

	for _, packagedChartPath := range packagedChartPaths {
		chartURL := filepath.Join(packageOutputDir, filepath.Base(packagedChartPath))
		if err := addToIndexFile(index, packagedChartPath, chartURL); err != nil {
			log.Fatalf("Failed to add chart to index file: %v", err)
		}

	}

	if err := saveIndexFile(index, indexFilePath); err != nil {
		log.Fatalf("Failed to save index file: %v", err)
	}

	//uploadToJfrogArtifactory()
}

func processChart(chartName string) {
	chartPath := filepath.Join(chartBasePath, chartName)
	log.Infof("Loading chart from %s", chartPath)

	chart, err := loadingChart(chartPath)
	if err != nil {
		log.Errorf("Error loading chart: %v", err)
		return
	}

	if err := updateAndBuildDependencies(chartPath); err != nil {
		log.Errorf("Error updating dependencies: %v", err)
		return
	}

	if err := lintChart(chart); err != nil {
		log.Errorf("Linting Errors found in chart %s: %v", chart.Name(), err)
		return
	}

	packagedChartPath, err := packageChart(chartPath, packageOutputDir)
	if err != nil {
		log.Fatalf("Failed to package chart: %v", err)
	}
	log.Info(packagedChartPath)

	packagedChartPaths = append(packagedChartPaths, packagedChartPath)
}

func loadingChart(charPath string) (*chart.Chart, error) {
	return loader.Load(charPath)
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

func updateAndBuildDependencies(chartPath string) error {
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

func downloadIndexFile() error {
	jfrogUploadAPIKey := os.Getenv("jfrog_upload_api_key")
	JFROGindexFileUrl := jfrogArtifactURL + jfrogRepoName + "/index.yaml"
	log.Infof("Index file URL is %s", JFROGindexFileUrl)

	request, err := http.NewRequest("GET", JFROGindexFileUrl, nil)
	if err != nil {
		return fmt.Errorf("Request is failed %v", err)
	}
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("Authorization", "Bearer "+jfrogUploadAPIKey)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return fmt.Errorf("Failed to download Index file %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Error while download index.yaml, Status Error Code: %d", response.StatusCode)
	}

	if err := os.MkdirAll(filepath.Dir(indexFilePath), os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory for index file %v", err)
	}

	out, err := os.Create(indexFilePath)

	if err != nil {
		return fmt.Errorf("failed to create index file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}

func packageChart(chartPath, outputDir string) (string, error) {
	chart, err := loader.Load(chartPath)
	if err != nil {
		return "", fmt.Errorf("failed to load chart: %w", err)
	}

	packagedChartPath, err := chartutil.Save(chart, outputDir)
	if err != nil {
		return "", fmt.Errorf("failed to package chart: %w", err)
	}

	log.Printf("Chart %s packaged successfully at %s", chart.Metadata.Name, packagedChartPath)
	return packagedChartPath, nil
}

func loadOrCreateIndex(indexPath string) (*repo.IndexFile, error) {
	index := repo.NewIndexFile()

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		log.Println("No existing index file found, creating a new index")
		return index, nil
	}

	data, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %w", err)
	}

	if err := yaml.Unmarshal(data, index); err != nil {
		return nil, fmt.Errorf("failed to unmarshal index file: %w", err)
	}

	log.Println("Loaded existing index file")
	return index, nil
}

func saveIndexFile(index *repo.IndexFile, path string) error {
	index.SortEntries()
	data, err := yaml.Marshal(index)
	if err != nil {
		return fmt.Errorf("failed to marshal index file: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write index file: %w", err)
	}

	log.Printf("Index file saved at %s", path)
	return nil
}

func uploadToJfrogArtifactory() {
	jfrogUploadAPIKey := os.Getenv("jfrog_upload_api_key") // I have exported key in OS already in form of env variable
	entries, err := os.ReadDir(packageOutputDir)
	if err != nil {
		log.Errorf("Reading temp directory of packaged charts has problem %v ", err)
	}

	for _, entry := range entries {
		chartPath := filepath.Join(packageOutputDir, entry.Name())
		file, err := os.Open(chartPath)
		if err != nil {
			log.Errorf("Error whilst reading %v", err)
		}
		defer file.Close()

		uploadURL := fmt.Sprintf("%s%s/%s", jfrogArtifactURL, jfrogRepoName, filepath.Base(chartPath))
		request, err := http.NewRequest("PUT", uploadURL, file)
		if err != nil {
			log.Errorf("Error whilst creating request %v", err)
		}

		request.Header.Set("Content-Type", "application/octet-stream")
		request.Header.Set("Authorization", "Bearer "+jfrogUploadAPIKey)

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Errorf("error during request: %v", err)
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
			log.Errorf("failed to upload file: %s", response.Status)
		}

		fmt.Printf("Uploaded %s successfully to %s\n", chartPath, uploadURL)
	}
}

func addToIndexFile(index *repo.IndexFile, chartPath, url string) error {
	chart, err := loader.Load(chartPath)
	if err != nil {
		return fmt.Errorf("failed to load chart: %w", err)
	}

	chartDetails := &chartDetails{
		APIVersion:  chart.Metadata.APIVersion,
		Created:     time.Now().Format(time.RFC3339Nano),
		Description: chart.Metadata.Description,
		Digest:      "", // I will add it later
		Home:        chart.Metadata.Home,
		Keywords:    chart.Metadata.Keywords,
		Maintainers: chart.Metadata.Maintainers,
		Name:        chart.Metadata.Name,
		Sources:     chart.Metadata.Sources,
		URLs:        []string{url},
		AppVersion:  chart.Metadata.AppVersion,
		Version:     chart.Metadata.Version,
	}

	fmt.Println(chartDetails)

	version := &repo.ChartVersion{
		Metadata: chart.Metadata,
		URLs:     []string{url},
		Created:  time.Now(),
	}

	if index.Entries[chart.Metadata.Name] == nil {
		index.Entries[chart.Metadata.Name] = repo.ChartVersions{}
	}

	index.Entries[chart.Metadata.Name] = append(index.Entries[chart.Metadata.Name], version)
	log.Printf("Added chart %s version %s to index", chart.Metadata.Name, chart.Metadata.Version)
	return nil
}
