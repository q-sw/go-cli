package terraform

import (
    "fmt"
    "os"
    "slices"
    "strings"

    "path/filepath"

    "github.com/q-sw/go-cli-qsw/internal/utils"
)

var PROJECTTYPE = []string{"standard", "env"}

var PROJECTFILES = []string{"providers.tf", "main.tf",
    "variables.tf", "versions.tf",
    "backend.tf", "terraform.tfvars", "outputs.tf", "README.md"}

var ENVPROJECTFILES = []string{"providers.tf", "main.tf",
    "variables.tf", "versions.tf", "outputs.tf", "README.md"}

var ENVFILES = []string{"backend.tf", "terraform.tfvars"}

var MODULEFILES = []string{"main.tf",
    "variables.tf", "versions.tf", "outputs.tf", "README.md"}

func Tf(instanceType, instanceName, instancePath, projectType string, enableGit bool) error {

    switch {
    case instanceType == "module":
        createModule(instanceName, instancePath, enableGit)
    case instanceType == "project":
        if slices.Contains(PROJECTTYPE, strings.ToLower(projectType)) == false {
            return fmt.Errorf("project type %v doesn't exists", projectType)
        }
        createProject(instanceName, instancePath, projectType, enableGit)
    case instanceType == "files":
        for _, f := range PROJECTFILES {
            utils.CreateFile(instancePath, f)
        }
        if enableGit {
            utils.GitInit(instancePath)
        }
    }
    return nil
}

func createModule(name, path string, git bool) error {
    var modulePath string
    if path == "" {
        currentPath, err := os.Getwd()
        if err != nil {
            return err
        }
        modulePath = filepath.Join(currentPath, "modules", name)
    }
    modulePath = filepath.Join(path, "modules", name)

    os.MkdirAll(modulePath, 0755)

    for _, f := range MODULEFILES {
        utils.CreateFile(modulePath, f)
    }
    if git {
        utils.GitInit(modulePath)
    }
    return nil
}

func createProject(name, path, projectType string, git bool) error {
    var projectPath string
    if path == "" {

        currentPath, err := os.Getwd()
        if err != nil {
            return err
        }
        projectPath = filepath.Join(currentPath, name)
    }
    projectPath = filepath.Join(path, name)

    os.MkdirAll(projectPath, 0755)
    os.MkdirAll(filepath.Join(projectPath, "modules"), 0755)

    if projectType == "env" {
        prodPath := filepath.Join(projectPath, "prod")
        devPath := filepath.Join(projectPath, "dev")

        os.MkdirAll(prodPath, 0755)
        os.MkdirAll(devPath, 0755)

        for _, f := range ENVPROJECTFILES {
            utils.CreateFile(projectPath, f)
        }

        for _, f := range ENVFILES {
            utils.CreateFile(prodPath, f)
            utils.CreateFile(devPath, f)
        }
    } else {

        for _, f := range PROJECTFILES {
            utils.CreateFile(projectPath, f)
        }
    }

    if git {
        utils.GitInit(projectPath)
    }

    return nil
}

// TODO: add tempating fonction to create README.md, providers.tf, versions.tf, backend.tf makefile with a default structure
