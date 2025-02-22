/*
Copyright © 2024 q-sw
*/
package cmd

import (
    "github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
    Use:   "get",
    Short: "Get information from your system or projet",
}

func init() {
    rootCmd.AddCommand(getCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // getCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
