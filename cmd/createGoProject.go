/*
Copyright © 2024 q-sw
*/
package cmd

import (
	"fmt"

	"github.com/q-sw/go-cli-qsw/internal/golang"
	"github.com/spf13/cobra"
)

// createGoProjectCmd represents the createGoProject command
var createGoProjectCmd = &cobra.Command{
	Use:     "go-project",
	Example: "cli create go-project [flags]",
	Short:   "create skaffold for Go project",
	Long:    `create skaffold for Go project like standard package, API, cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createGoProject called")
		golang.Creator(goProjectType)
	},
}

var goProjectType string
var goEnableGit bool

func init() {
	createCmd.AddCommand(createGoProjectCmd)

	createGoProjectCmd.Flags().StringVarP(&goProjectType, "type", "t", "package", "[Global] Go project type standard|api|cli")
	createGoProjectCmd.Flags().BoolVarP(&goEnableGit, "enable-git", "g", false, "[Global] Activate Git ")
}
