package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("file", "mylinks.json", "JSON file containing a list of URLs to screenshot.")
	outputDir := flag.String("output", "website-screenshots", "Directory to save screenshots.")
	flag.Parse()

	// Read the JSON file
	jsonFile, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Error opening JSON file %s: %v", *inputFile, err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var urls []string
	json.Unmarshal(byteValue, &urls)

	// Create the output directory if it doesn't exist
	if _, err := os.Stat(*outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(*outputDir, 0755)
		if err != nil {
			log.Fatalf("Error creating output directory %s: %v", *outputDir, err)
		}
	}

	// Loop through URLs and take screenshots
	for _, rawURL := range urls {
		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			log.Printf("Skipping invalid URL: %s", rawURL)
			continue
		}

		// Generate a filename from the URL
		filename := fmt.Sprintf("%s.png", parsedURL.Host+parsedURL.Path)
		// Replace slashes in path with underscores for a valid filename
		filename = filepath.Join(*outputDir, SanitizeFileName(filename))

		fmt.Printf("Taking screenshot of %s and saving to %s\n", rawURL, filename)

		// Execute gowitness command
		cmd := exec.Command("gowitness", "single", "-o", filename, rawURL)
		err = cmd.Run()
		if err != nil {
			log.Printf("Error taking screenshot for %s: %v", rawURL, err)
		}
	}

	fmt.Println("Screenshot process completed.")
}

// SanitizeFileName replaces characters that are invalid in file names.
func SanitizeFileName(name string) string {
	// Replace path separators with a different character
	return url.PathEscape(name)
}
