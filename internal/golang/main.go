package golang

import (
	"os"
	"path/filepath"

	"github.com/q-sw/go-cli-qsw/internal/utils"
)

var goPackageFolder = []string{"cmd", "internal", "pkg", "bin", "test"}
var goAPIFolder = []string{"api", "storage", "types", "utils", "test", "bin"}
var goStdFiles = []string{"README.md", "LICENSE", "makefile", "main.go"}

func Creator(actionType string) {
	currentDir, _ := os.Getwd()
	switch {
	case actionType == "package":
		for _, f := range goStdFiles {
			utils.CreateFile(currentDir, f)
		}
		for _, d := range goPackageFolder {
			os.MkdirAll(filepath.Join(currentDir, d), 0755)
		}
	case actionType == "api":
		for _, f := range goStdFiles {
			utils.CreateFile(currentDir, f)
		}
		for _, d := range goAPIFolder {
			os.MkdirAll(filepath.Join(currentDir, d), 0755)
		}
	}

}

// TODO: add project with prefix like github.com/q-sw/{{project_name}}
// TODO: use config file to configure the project prefix
// TODO: add template for README.md makefile, main.go
