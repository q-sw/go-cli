package golang

import (
	"log"
	"os"
	"os/exec"

	"github.com/q-sw/go-qsw-cli/pkg/utils"
)

var goPackageFolder = []string{"cmd", "internal", "pkg", "bin", "test"}
var goAPIFolder = []string{"api", "storage", "types", "utils", "test", "bin"}
var goStdFiles = []string{"README.md", "LICENSE", "makefile", "main.go"}

func Creator(actionType string, projectName string) {
	currentDir, _ := os.Getwd()
	switch {
	case actionType == "package":
		for _, f := range goStdFiles {
			utils.CreateFile(utils.File{f, currentDir})
		}
		for _, d := range goPackageFolder {
			utils.CreateDir(utils.Directory{d, currentDir})
		}
	case actionType == "api":
		for _, f := range goStdFiles {
			utils.CreateFile(utils.File{f, currentDir})
		}
		for _, d := range goAPIFolder {
			utils.CreateDir(utils.Directory{d, currentDir})
		}
	}

	cmd := exec.Command("go", "mod", "init", projectName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
