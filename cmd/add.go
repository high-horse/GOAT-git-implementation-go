package cmd

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
	_"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add [files...]",
	Short: "Stage files for commit",
	Long:  "The add command stages the specified files, preparing them to be committed.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: No files specified")
			return
		}

		for _, arg := range args {
			if arg == "." {
				if err := addAllFiles("."); err != nil {
					fmt.Printf("Error adding files: %v\n", err)
				}
			} else {
				if err := addFile(arg); err != nil {
					fmt.Printf("Error adding file %s: %v\n", arg, err)
				}
			}
		}
	},
}

const goatDir = ".goat"

func addFile(filePath string) error {
	if !fileExists(filePath) {
		return fmt.Errorf("file %s does not exist", filePath)
	}

	// Calculate SHA-1 hash
	hash, err := calculateFileHash(filePath)
	if err != nil {
		return fmt.Errorf("calculating hash for file %s: %w", filePath, err)
	}

	// Save file content to object store
	objectFilePath := filepath.Join(goatDir, "objects", hash[:2], hash[2:])
	if err := saveFileToObjectStore(filePath, objectFilePath); err != nil {
		return fmt.Errorf("saving file %s to object store: %w", filePath, err)
	}

	// Update the index
	if err := updateIndex(filePath, hash); err != nil {
		return fmt.Errorf("updating index for file %s: %w", filePath, err)
	}

	fmt.Printf("File %s added successfully\n", filePath)
	return nil
}

func addAllFiles(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories, hidden files, and the .goat directory
		if info.IsDir() {
            if info.Name() == ".git" || info.Name() == ".goat" {
                return filepath.SkipDir
            }
            return nil
        }
		// if info.IsDir() || strings.HasPrefix(info.Name(), ".") || path == goatDir {
		// 	return nil
		// }
		return addFile(path)
	})
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func saveFileToObjectStore(filePath, objectFilePath string) error {
	if err := os.MkdirAll(filepath.Dir(objectFilePath), 0755); err != nil {
		return err
	}

	return os.Link(filePath, objectFilePath)
}

func updateIndex(filePath, hash string) error {
	indexFilePath := filepath.Join(goatDir, "index")
	indexFile, err := os.OpenFile(indexFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer indexFile.Close()

	_, err = fmt.Fprintf(indexFile, "%s %s\n", hash, filePath)
	return err
}