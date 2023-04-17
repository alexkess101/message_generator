package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	parentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	zshrcFile := filepath.Join(os.Getenv("HOME"), ".zshrc")
	f, err := os.OpenFile(zshrcFile, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("\n\n#---------------------------MESSAGE ALIASES---------------------------\n"))
	if err != nil {
		fmt.Printf("Failed to write to .zshrc file: %v", err)
	}

	err = filepath.Walk(parentDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.IsDir() && info.Name() != parentDir {
			mainFile := filepath.Join(path, "main.go")
			if _, err := os.Stat(mainFile); err == nil {
				dirName := filepath.Base(path)

				aliasName := fmt.Sprintf("%s", dirName)
				aliasCmd := fmt.Sprintf("alias %s='go run %s/main.go \"$@\"'\n", aliasName, path)

				if err != nil {
					return fmt.Errorf("failed to open .zshrc file: %v", err)
				}
				if _, err = f.WriteString(aliasCmd); err != nil {
					return fmt.Errorf("failed to write alias command to .zshrc file: %v", err)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
