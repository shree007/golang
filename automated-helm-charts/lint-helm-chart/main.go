package main

import (
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/lint"
	"helm.sh/helm/v3/pkg/lint/support"
)

/*
I need to print chart name as well like
Chart Name Linting completed with 1 messages:
[INFO] icon is recommended
*/

func main() {
	chartPath := "../helm/charts/prometheus-adapter"
	namespace := "default"
	strictMode := true

	linter := lint.All(chartPath, map[string]interface{}{}, namespace, strictMode)
	if len(linter.Messages) > 0 {
		fmt.Printf("Linting completed with %d messages:\n", len(linter.Messages))
		for _, msg := range linter.Messages {
			fmt.Printf("[%s] %s\n", severityToString(msg.Severity), msg.Err)
		}
	} else {
		fmt.Println("No linting issues found!")
	}

	if linter.HighestSeverity >= support.ErrorSev {
		fmt.Println("Errors found during linting!")
		os.Exit(1)
	}
}

func severityToString(sev int) string {
	switch sev {
	case support.ErrorSev:
		return "ERROR"
	case support.WarningSev:
		return "WARNING"
	case support.InfoSev:
		return "INFO"
	default:
		return "UNKNOWN"
	}
}
