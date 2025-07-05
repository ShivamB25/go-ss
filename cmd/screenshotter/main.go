package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"os/exec"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("file", "mylinks.json", "JSON file containing a list of URLs to screenshot.")
	outputDir := flag.String("output", "website-screenshots", "Directory to save screenshots.")
	flag.Parse()

	urls, err := loadURLs(*inputFile)
	if err != nil {
		log.Fatalf("Error loading URLs from %s: %v", *inputFile, err)
	}

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory %s: %v", *outputDir, err)
	}

	for _, rawURL := range urls {
		if err := takeScreenshot(rawURL, *outputDir); err != nil {
			log.Printf("Failed to take screenshot for %s: %v", rawURL, err)
		}
	}

	fmt.Println("Screenshot process completed.")
}

func loadURLs(filename string) ([]string, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var urls []string
	if err := json.Unmarshal(byteValue, &urls); err != nil {
		return nil, fmt.Errorf("unmarshaling json: %w", err)
	}

	return urls, nil
}

func takeScreenshot(rawURL, outputDir string) error {
	// Validate URL before proceeding
	if _, err := url.ParseRequestURI(rawURL); err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	fmt.Printf("Taking screenshot of %s (saving to %s)\n", rawURL, outputDir)

	// Execute gowitness command using the --destination flag
	cmd := exec.Command("gowitness", "single", "--destination", outputDir, rawURL)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("running gowitness for %s: %w\nOutput: %s", rawURL, err, string(output))
	}

	return nil
}
