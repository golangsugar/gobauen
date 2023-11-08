package gobauen

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"regexp"
)

const (
	paramModel        = 0
	paramProjectName  = 1
	paramDirectory    = 2
	paramsLengthMin   = 3
	modelsFileDefault = "gobauen.yml"
)

var rxDirectory = regexp.MustCompile(`^[\p{L}\p{N}_-]+$`)

func ParseCommand() (modelName, projectName, targetDirectory string, err error) {
	if len(os.Args) < paramsLengthMin {
		printUsage()
		return "", "", "", ErrInsufficientParams
	}

	modelName = os.Args[paramModel]
	projectName = os.Args[paramProjectName]
	targetDirectory = os.Args[paramDirectory]

	if projectName == "" {
		return "", "", "", ErrUndefinedProject
	}

	if targetDirectory == "" {
		targetDirectory = projectName
	}

	if !rxDirectory.MatchString(targetDirectory) {
		return "", "", "", fmt.Errorf("invalid directory name: %q", targetDirectory)
	}

	return modelName, projectName, targetDirectory, nil
}

// Model returns the model with the given name
func Model(modelName string) (ProjectModel, error) {
	modelsFile := flag.String("f", modelsFileDefault, fmt.Sprintf("models file. defaults to %q", modelsFileDefault))
	flag.Parse()

	buf, err := os.ReadFile(*modelsFile)
	if err != nil {
		return ProjectModel{}, fmt.Errorf("error loading models from file %q: %w", modelsFile, err)
	}

	var a []ProjectModel
	if err = yaml.Unmarshal(buf, a); err != nil {
		return ProjectModel{}, fmt.Errorf("error decoding model from file %q: %w", modelsFile, err)
	}

	for _, model := range a {
		if model.Name == modelName {
			return model, nil
		}
	}

	return ProjectModel{}, fmt.Errorf("unknown model %q requested", modelName)
}

func Perform(model ProjectModel, projectName, directory string) error {
	cmd := exec.Command("git", "clone", "--depth=1", model.Repository, directory)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error cloning repository %q: %w", model.Repository, err)
	}

	if err := os.RemoveAll(fmt.Sprintf("%s/.git", directory)); err != nil {
		return fmt.Errorf("error removing .git directory: %w", err)
	}

	for _, model := range md {
		if model.Name == projectName {
			cmd := exec.Command("git", "clone", model.Repository, directory)
			out, err := cmd.CombinedOutput()
		}
	}

	return fmt.Errorf("undefined model: %q")
}

var (
	ErrInsufficientParams = fmt.Errorf("insufficient params")
	ErrUndefinedProject   = fmt.Errorf("please, provide the project name")
	ErrEmptyDirectives    = fmt.Errorf("empty modelDirectives")
)

func ptrValue[T any](v *T) T {
	if v == nil {
		return *new(T)
	}

	return *v
}

type executionDirectives struct {
	projectName,
	directory string
	md ProjectModel
}

func getModelDirectives(modelName string) (*ProjectModel, error) {
	if len(os.Args) < paramsLengthMin {
		printUsage()
		return nil, ErrInsufficientParams
	}

	buf, err := os.ReadFile(directivesFile)
	if err != nil {
		return nil, fmt.Errorf("error loading modelDirectives from file %q: %w", directivesFile, err)
	}

	type models struct {
		Models []ProjectModel `yaml:"models"`
	}

	var a models
	if err = yaml.Unmarshal(buf, &a); err != nil {
		return nil, fmt.Errorf("error decoding modelDirectives from file %q: %w", directivesFile, err)
	}

	for _, model := range a.Models {
		if model.Name == modelName {
			return &model, nil
		}
	}

	return nil, fmt.Errorf("unknown model %q", modelName)
}
