package main

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
)

func main() {
	chartBasePath := "helm/charts"
	log.Info("Read Base Directory: ", chartBasePath)
	entries, err := os.ReadDir(chartBasePath)

	if err != nil {
		log.Error("Error Reading chart directory", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			chartPath := filepath.Join(chartBasePath, entry.Name())
			log.Infof("Loading chart from %s", chartPath)
			chart, err := loader.Load(chartPath)

			if err != nil {
				log.Error("Erros in loading charts: ", err)
				continue
			}
			log.Infof("Loaded charts are %s", chart.Name())
			settings := cli.New()
			manager := &downloader.Manager{
				ChartPath:  chartPath,
				Getters:    getter.All(settings),
				SkipUpdate: false,
				Out:        os.Stdout,
			}

			if err := manager.Update(); err != nil {
				log.Errorf("Error while updating dependencies: %v", err)
			} else {
				log.Infof("Dependencies updated successfully")
			}

			if err := manager.Build(); err != nil {
				log.Errorf("Error while building dependencies in chart %s: %v", chart.Name(), err)
				continue
			}
			log.Infof("Dependencies built successfully for chart %s", chart.Name())

			if err := lintChart(chart); err != nil {
				log.Errorf("Linting Errors found in chart %s: %v", chart.Name(), err)
			} else {
				log.Infof("Linting passed for chart %s", chart.Name())
			}

			packagePath := "temp-helm-storage"

			pkgPath, err := chartutil.Save(chart, packagePath)
			if err != nil {
				log.Errorf("error saving packaged chart %s: %v", chart.Name(), err)
			}
			log.Infof("Packaged chart %s to %s", chart.Name(), pkgPath)

		} else {
			log.Infof("%s is not a directory: ", entry.Name())
		}
	}
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
