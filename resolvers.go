package resolvers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	downloadURL  = "https://raw.githubusercontent.com/trickest/resolvers/main/resolvers.txt"
	folderName   = "Resolvers"
	fileName     = "resolvers.txt"
)

func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Create the Resolvers folder
	resolversPath := filepath.Join(currentDir, folderName)
	err = os.MkdirAll(resolversPath, 0755) // Use MkdirAll to create parent directories if they don't exist
	if err != nil {
		fmt.Println("Error creating Resolvers folder:", err)
		return
	}

	// Remove old files in the Resolvers folder
	files, err := filepath.Glob(filepath.Join(resolversPath, "*"))
	if err != nil {
		fmt.Println("Error getting list of old files:", err)
		return
	}

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Println("Error deleting old file:", err)
			return
		}
	}

	// Download the text file
	filePath := filepath.Join(resolversPath, fileName)
	err = downloadFile(filePath, downloadURL)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}

	fmt.Println("File downloaded successfully to:", filePath)
}

func downloadFile(filePath, url string) error {
	// Create the file to save the downloaded contents
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Perform the HTTP request to get the file contents
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Copy the contents of the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
