package cmd

import (
    "fmt"
    "io"
    "os"
    "os/user"
    "path/filepath"

    "github.com/spf13/cobra"
)

var usr, err = user.Current()
var sourceFile = filepath.Join(usr.HomeDir, "/cplib/template.cpp") 

var destinationPath string

var initCmd = &cobra.Command{
    Use: "init",
    Short: "Initializes a CPlib template in the specified location.",
    Run: func(cmd *cobra.Command, args[] string) {
        err := copyFile(sourceFile, destinationPath)
        if err != nil {
            fmt.Printf("Error Copying: %v \n", err)
            os.Exit(1)
        }
        fmt.Println("Template created.")
    },
}

func copyFile(src, dest string) error {
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()

    destination, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer destination.Close()

    _, err  = io.Copy(destination, source)
    return err
}

func init() {
	initCmd.Flags().StringVarP(&destinationPath, "destination", "d", "", "Destination path to create the template in.")

	initCmd.MarkFlagRequired("destination")
}
