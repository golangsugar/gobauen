package gobauen

// ProjectModel contains directives for generating a specific project-model directory tree
type ProjectModel struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}
