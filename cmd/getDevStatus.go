/*
Copyright © 2024 q-sw
*/
package cmd

import (
	"github.com/q-sw/go-cli-qsw/internal/devEnvStatus"
	"github.com/spf13/cobra"
)

// getDevStatusCmd represents the getDevStatus command
var getDevStatusCmd = &cobra.Command{
	Use:   "dev-status",
	Short: "show the status of your dev environments",
	Long: `show the status of your dev environments.
This command use the cli config file ${HOME}/.cli-config.yaml
`,
	Run: func(cmd *cobra.Command, args []string) {
		devenvstatus.GetDevStatus(statusVerbose)
	},
}

var statusVerbose bool
var showAllBranches bool
var showBranch bool

func init() {
	getCmd.AddCommand(getDevStatusCmd)
	getDevStatusCmd.Flags().BoolVarP(&statusVerbose, "verbose", "v", false, "[Global] Show details about repository status")
	getDevStatusCmd.Flags().BoolVarP(&showBranch, "show-branch", "b", false, "[Global] Show actual branch")
	getDevStatusCmd.Flags().BoolVarP(&showAllBranches, "show-all-branches", "", false, "[Global] Show local and remote branches")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getDevStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getDevStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
