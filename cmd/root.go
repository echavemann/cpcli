package cmd

import (

    "github.com/spf13/cobra"
)

// Root command definition and initialization

var rootCmd = &cobra.Command{
    Use: "cpg",
    Short: "Competitive Programming Utility",
    Run: func(cmd *cobra.Command, args[] string) {
    },
}

func init() {
	rootCmd.Flags().StringP("include", "i", "", "CPLib Packages to Include")

}

func Execute() error {
	return rootCmd.Execute()
}
