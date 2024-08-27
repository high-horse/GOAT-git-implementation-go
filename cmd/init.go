// cmd/init.go
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init()  {
	rootCmd.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use: "init",
	Short: "prints the status of repository",
	Long: "The status command shows the current status of the repository",
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
	},
}

func initRepo() {
	goatDir := ".goat"
	if err := os.Mkdir(goatDir, 0755); err != nil {
		fmt.Println("Error creating .goat directory:", err)
		return
	}
	
	dirs := []string{
		"objects", "refs",
	}
	for _, dir := range dirs {
		path := filepath.Join(goatDir, dir)
		if err := os.Mkdir(path, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	} 
	
	fmt.Println("Initialized empty directory in ", goatDir)
}