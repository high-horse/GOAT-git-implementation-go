// cmd/add.go
package cmd

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init()  {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command {
	Use: "add [files...]",
	Short: "stage files for commit",
	Long: "The add command stages the specified files, preparing them to be committed.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: No files specified")
			return
		}
		for _, file := range args {
			addFile(file)
		}
	},
}

var goatDir = "./goat"

func addFile(filePath string) {
	// goatDir := "./goat"
	if !fileExists(filePath) {
		fmt.Printf("Error: File %s does not exist\n", filePath)
		return
	}
	
	// calculate SHA-1 hash
	hash, err := calculateFileHash(filePath)
	if err != nil {
		fmt.Printf("Error calculating hash for file %s: %v\n", filePath, err)
		return
	}
	
	// Save file content to object store
	objectFilePath := filepath.Join(goatDir, "objects", hash)
	if err := saveFileToObjectStore(filePath, objectFilePath); err != nil {
		fmt.Printf("Error saving file %s: %v\n", filePath, err)
		return
	}
	
	// Update the index
	if err := updateIndex(filePath, hash); err != nil {
		fmt.Printf("Error updating index for file %s: %v\n", filePath, err)
		return
	}
	fmt.Printf("File %s added successfully\n", filePath)
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
	
	return fmt.Sprintf("%x",hash.Sum(nil)), nil
}

func saveFileToObjectStore(filePath, objectFilePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()
	
	objectFile, err := os.Create(objectFilePath)
	if err != nil {
		return err
	}
	defer objectFile.Close()
	
	if _, err := io.Copy(objectFile, file); err != nil {
		return err
	}
	
	return nil
}

func updateIndex(filePath, hash string) error {
	// goatDir
	indexFilePath := filepath.Join(goatDir, "index")
	indexFile, err := os.OpenFile(indexFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	
	defer indexFile.Close()
	_, err = fmt.Fprintf(indexFile, "%s %s\n", hash, filePath)
	return err
}