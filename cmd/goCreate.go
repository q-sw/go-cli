/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/q-sw/go-qsw-cli/pkg/golang"
	"github.com/q-sw/go-qsw-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// goCreateCmd represents the goCreate command
var goCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create some stuffs for Go",
	Long: `go create subcommand allow you to create project skaffold
or create simple all standard files you need to start with go`,
	ValidArgs: []string{"api", "package"},
	Args: func(cmd *cobra.Command, args []string) error {
		return utils.CheckArgs(cmd.ValidArgs, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		golang.Creator(args[0], projectName)
	},
}

func init() {
	goCmd.AddCommand(goCreateCmd)
	goCreateCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name for your project")
	goCreateCmd.Flags().BoolVarP(&enableGit, "enable-git", "g", false, "Activate Git (git init) on your folder/project")
	goCreateCmd.Flags().StringVarP(&projectPath, "path", "", "./", "Path where create file or project")
}
