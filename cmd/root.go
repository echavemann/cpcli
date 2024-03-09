package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

// Root command definition and initialization

var rootCmd = &cobra.Command{
    Use: "cpg",
    Short: "Competitive Programming Utility",
    Run: func(cmd *cobra.Command, args[] string) {
        fmt.Println("BAse Command Called!")
    },
}

func init() {
	rootCmd.Flags().StringP("include", "i", "", "CPLib Packages to Include")
    rootCmd.AddCommand(initCmd)

}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
