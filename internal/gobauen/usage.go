package gobauen

import (
	"fmt"
	"os"
	"path/filepath"
)

func printUsage() {
	bin := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %s projectname [output directory]\n", bin)
	fmt.Println("  projectname: name of the project to be generated")
	fmt.Println("  output directory: directory where the project will be generated. if empty, a subdirectory with\n the project name will be created in the current directory. if dot (.), the project will be generated in the current directory")
	fmt.Printf("\n\nExample: %s myproject .\n\n", bin)
}
