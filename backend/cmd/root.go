package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use: "arendt",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
