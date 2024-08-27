// cmd/init.go
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	if err := os.MkdirAll(goatDir, 0755); err != nil {
		fmt.Println("Error creating .goat directory:", err)
		return
	}
	
	// Ask the user for the remote repository URL
	fmt.Print("Enter the remote repository URL (leave empty to skip): ")
	reader := bufio.NewReader(os.Stdin)
	remoteURL, _ := reader.ReadString('\n')
	remoteURL = strings.TrimSpace(remoteURL)
	// If remoteURL is empty, assign the default path
	if remoteURL == "" {
		remoteURL = "/home/camle/Desktop"
		fmt.Println("No remote URL provided. Using default:", remoteURL)
	} else {
		fmt.Println("Using provided remote URL:", remoteURL)
	}

	
	dirs := []string{
		"hooks",
		"info",
		"objects", 
		"refs/heads",
		"refs/tags",
		"logs/refs/heads",
	}
	for _, dir := range dirs {
		path := filepath.Join(goatDir, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	} 
	
	// Prepare the config content
	configContent := `[core]
    repositoryformatversion = 0
    filemode = true
    bare = false
	`
	// if remoteURL != "" {
		configContent += fmt.Sprintf(`
[remote "origin"]
    url = %s
    fetch = +refs/heads/*:refs/remotes/origin/*
`, remoteURL)
	// }
	fmt.Println("remoteurl", remoteURL)

	
	files := map[string]string{
		"HEAD": "ref: refs/heads/main\n",
		"config": configContent,
		"description": "Unnamed repository; edit this file 'description' to name the repository.\n",
		"info/exclude": "# git ls-files --others --exclude-from=.git/info/exclude\n" +
			"# Lines that start with '#' are comments.\n" +
			"# For a project mostly in C, the following would be a good set of\n" +
			"# exclude patterns (uncomment them if you want to use them):\n" +
			"# *.[oa]\n" +
			"# *~\n",
	}
	
	for file, content :=  range files {
		path := filepath.Join(goatDir, file) 
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	}
	fmt.Println("Initialized empty directory in ", goatDir)
}