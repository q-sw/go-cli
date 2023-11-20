package terraform

import (
	"os"
	"path/filepath"

	"github.com/q-sw/go-qsw-cli/pkg/utils"
)

var terraformFiles = []string{"providers.tf", "main.tf",
	"variables.tf", "versions.tf",
	"backend.tf", "terraform.tfvars", "outputs.tf", "README.md"}

var terraformModuleFiles = []string{"main.tf",
	"variables.tf", "versions.tf", "outputs.tf", "README.md"}

func createStdProject(projectName string, projectPath string) {
	utils.CreateDir(utils.Directory{projectName, projectPath})
	p := filepath.Join(projectPath, projectName)
	utils.CreateDir(utils.Directory{"modules", p})
	createFilesOnly(terraformFiles, p)
	utils.InitRepo(p, terraformFiles)
}

func createModule(moduleName string) {
	currentDir, _ := os.Getwd()
	modulePath := filepath.Join(currentDir, "modules")

	mp := utils.Directory{moduleName, modulePath}
	utils.CreateDir(mp)

	createFilesOnly(terraformModuleFiles, filepath.Join(modulePath, moduleName))
}

func createFilesOnly(files []string, filePath string) {
	var tfFile []utils.File

	for _, f := range files {
		tfFile = append(tfFile, utils.File{f, filePath})
	}

	utils.CreateFiles(tfFile)

}

func Creator(actionType string, projectName string, projectType string, projectPath string) {

	switch {
	case actionType == "project" && projectType == "standard":
		createStdProject(projectName, projectPath)
	case actionType == "module":
		createModule(projectName)
	}
}
