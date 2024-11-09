package main

import (
	"fmt"
	"io"
	"net/http"
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
	chartBasePath    = "helm/charts"
	packagePath      = "temp-helm-storage"
	indexFilePath    = "temp-helm-storage/index.yaml"
	jfrogArtifactURL = "https://linkinpark.jfrog.io/artifactory/"
	jfrogRepoName    = "linkinpark-helmchart-helm-local"
)

/*
Cleanup index file old packaged charts and before starting
*/
func init() {
	if _, err := os.Stat(packagePath); os.IsNotExist(err) {
		log.Errorf("Directory does not exists, nothing to remove %v", err)
	}

	if err := os.RemoveAll(packagePath); err != nil {
		log.Errorf("Failed to remove directory %v", err)
	} else {
		log.Info("Directory has been removed")
	}
}

func main() {
	log.Info("Starting Helm chart processing")

	if err := downloadIndexFile(); err != nil {
		log.Fatalf("Failed to Download Index file %v", err)
	}

	index := manageIndexFile(indexFilePath)

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
	uploadToJfrogArtifactory()

}

func processChart(chartName string, index *repo.IndexFile) {
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

	chartURL := packageChart(chart)
	addToIndex(chart, chartURL, index)
}

func loadingChart(charPath string) (*chart.Chart, error) {
	return loader.Load(charPath)
}

func manageIndexFile(indexFilePath string) *repo.IndexFile {
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
	chartName := chart.Metadata.Name
	chartVersion := chart.Metadata.Version
	if existingVersions, ok := index.Entries[chart.Metadata.Name]; ok {
		for _, v := range existingVersions {
			if v.Version == chart.Metadata.Version {
				break
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
	return fmt.Sprintf("%s%s-%s.tgz", packagePath, chart.Metadata.Name, chart.Metadata.Version)
}

func uploadToJfrogArtifactory() {
	jfrogUploadAPIKey := os.Getenv("jfrog_upload_api_key") // I have exported key in OS already in form of env variable
	entries, err := os.ReadDir(packagePath)
	if err != nil {
		log.Errorf("Reading temp directory of packaged charts has problem %v ", err)
	}

	for _, entry := range entries {
		chartPath := filepath.Join(packagePath, entry.Name())
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
			log.Errorf("error during request: %w", err)
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
			log.Errorf("failed to upload file: %s", response.Status)
		}

		fmt.Printf("Uploaded %s successfully to %s\n", chartPath, uploadURL)
	}
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
