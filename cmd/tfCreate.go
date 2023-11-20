/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/q-sw/go-qsw-cli/pkg/terraform"
	"github.com/q-sw/go-qsw-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// tfCreateCmd represents the tfCreate command
var tfCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create some stuffs for Terraform",
	Long: `tf create subcommand allow you to create project/module skaffold
or create simple all standard files you need to start with terraform`,
	ValidArgs: []string{"project", "module", "files"},
	Args: func(cmd *cobra.Command, args []string) error {
		return utils.CheckArgs(cmd.ValidArgs, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		terraform.Creator(args[0], projectName, projectType, projectPath)
	},
}

var projectName string
var projectType string
var projectPath string
var enableGit bool

func init() {
	tfCmd.AddCommand(tfCreateCmd)
	tfCreateCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name for your project")
	tfCreateCmd.Flags().StringVarP(&projectType, "type", "t", "", "Project type, accept standard env platform")
	tfCreateCmd.Flags().BoolVarP(&enableGit, "enable-git", "g", false, "Activate Git (git init) on your folder/project")
	tfCreateCmd.Flags().StringVarP(&projectPath, "path", "", "./", "Path where create file or project")

}

func myUsageFunc(cobra.Command) {}
