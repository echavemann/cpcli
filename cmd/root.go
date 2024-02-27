package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// Root command definition and initialization

var rootCmd = &cobra.Command{
    Use: "cpg",
    Short: "Competitive Programming Utility",
    Run: func(cmd *cobra.Command, args[] string) {
        source, _ := cmd.Flags().GetString("source")
        destination, _ := cmd.Flags().GetString("destination")

        if err := copyFileOrDir(source, destination); err != nil {
            fmt.Printf("Error")
        }
    },
}

func init() {
	rootCmd.Flags().StringP("source", "s", "", "Source file or directory path")
	rootCmd.Flags().StringP("destination", "d", "", "Destination directory path")

	rootCmd.MarkFlagRequired("source")
	rootCmd.MarkFlagRequired("destination")
}

func Execute() error {
	return rootCmd.Execute()
}
