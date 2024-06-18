package golang

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/q-sw/go-cli-qsw/internal/templates"
	"github.com/q-sw/go-cli-qsw/internal/utils"
	"github.com/spf13/viper"
)

var goPackageFolder = []string{"cmd", "internal", "bin", "test"}
var goAPIFolder = []string{"api", "storage", "types", "utils", "test", "bin"}
var goStdFiles = []string{"README.md", "LICENSE", "makefile", "main.go"}

func Creator(actionType, name, description string) {
	if name != "" {
		os.Mkdir(name, 0755)
		os.Chdir(name)
	}
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
	templates.RenderReadme(name, description)
	templates.RenderMakefile(name)
	templates.RenderMain(name)
	initGoProject(name)
}

func initGoProject(name string) {
	if name == "" {
		name = utils.GetCurrentDirName()
	}
	p := viper.GetString("github_profile")
	prefix := p + "/" + name

	cmd := exec.Command("go", "mod", "init", prefix)
	cmd.Run()
}
