package main

import (
	"goko/internal/gobauen"
	"log/slog"
	"os"
)

func main() {
	modelName, projectName, targetDirectory, err := gobauen.ParseCommand()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	model, err := gobauen.Model(modelName)
	if err != nil {
		return "", "", err
	}

	allModelDirectives, err := getModelDirectives(projectName)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err = perform(projectName, targetDirectory, allModelDirectives); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
