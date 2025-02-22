/*
Copyright © 2024 q-sw
*/
package cmd

import (
    "github.com/q-sw/go-cli/internal/terraform"
    "github.com/spf13/cobra"
)

// createTerraformModuleCmd represents the createTerraformModule command
var createTerraformCmd = &cobra.Command{
    Use:       "terraform",
    Example:   "cli create terrafrom [project|module|files] [flags]",
    Short:     "create skaffold for terraform",
    Long:      `create skaffold for terraform project and module`,
    ValidArgs: []string{"project", "module", "files"},
    Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
    Run: func(cmd *cobra.Command, args []string) {
        terraform.Tf(args[0], instanceName, instancePath, projectType, enableGit)
    },
}

var instanceName string
var instancePath string
var enableGit bool
var projectType string

func init() {
    createCmd.AddCommand(createTerraformCmd)
    createTerraformCmd.Flags().StringVarP(&instanceName, "name", "n", "", "[Global] Name for your terraform module or project")
    createTerraformCmd.Flags().StringVarP(&projectType, "type", "t", "standard", "[Project] Project type, accept standard;env")
    createTerraformCmd.Flags().BoolVarP(&enableGit, "enable-git", "g", false, "[Project] Activate Git (git init) on your folder/project")
    createTerraformCmd.Flags().StringVarP(&instancePath, "path", "p", "./", "[Global] Path where create your instance")
}
