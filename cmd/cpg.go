package cmd

import (
	"io"
	"os"
	"path/filepath"
)


// Command Implementations

func copyFileOrDir(source, destination string) error {
	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return copyDir(source, destination)
	} else {
		return copyFile(source, destination)
	}
}

func copyDir(source, destination string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, _ := filepath.Rel(source, path)
		destPath := filepath.Join(destination, relativePath)

		if info.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		} else {
			return copyFile(path, destPath)
		}
	})
}

func copyFile(source, destination string) error {
    sourceFile, err := os.Open(source)
    if err != nil {
		return err
	}
    defer sourceFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
    defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
