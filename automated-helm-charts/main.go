package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/chart/loader"
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
		} else {
			log.Infof("%s is not a directory: ", entry.Name())
		}
	}
}
