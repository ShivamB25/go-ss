package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/playwright-community/playwright-go"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("file", "mylinks.json", "JSON file containing a list of URLs to screenshot.")
	outputDir := flag.String("output", "website-screenshots", "Directory to save screenshots.")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent screenshot operations.")
	flag.Parse()

	err := playwright.Install()
	if err != nil {
		log.Fatalf("could not install playwright dependencies: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	defer func() {
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop playwright: %v", err)
		}
	}()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch chromium: %v", err)
	}
	defer func() {
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
	}()

	urls, err := loadURLs(*inputFile)
	if err != nil {
		log.Fatalf("Error loading URLs from %s: %v", *inputFile, err)
	}

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory %s: %v", *outputDir, err)
	}

	var wg sync.WaitGroup
	urlChan := make(chan string)

	// Start worker goroutines
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rawURL := range urlChan {
				if err := takeScreenshot(browser, rawURL, *outputDir); err != nil {
					log.Printf("Failed to take screenshot for %s: %v", rawURL, err)
				}
			}
		}()
	}

	// Feed URLs to the workers
	for _, u := range urls {
		urlChan <- u
	}
	close(urlChan)

	wg.Wait()
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

func takeScreenshot(browser playwright.Browser, rawURL, outputDir string) error {
	page, err := browser.NewPage()
	if err != nil {
		return fmt.Errorf("could not create page: %w", err)
	}
	defer page.Close()

	fmt.Printf("Navigating to %s\n", rawURL)
	_, err = page.Goto(rawURL, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		Timeout:   playwright.Float(60000), // 60 seconds
	})
	if err != nil {
		return fmt.Errorf("could not navigate to %s: %w", rawURL, err)
	}

	safeFilename := SanitizeFileName(rawURL) + ".png"
	outputPath := filepath.Join(outputDir, safeFilename)

	fmt.Printf("Taking screenshot of %s -> %s\n", rawURL, outputPath)
	if _, err := page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String(outputPath),
		FullPage: playwright.Bool(true),
	}); err != nil {
		return fmt.Errorf("could not take screenshot for %s: %w", rawURL, err)
	}

	return nil
}

// SanitizeFileName creates a safe filename from a URL
func SanitizeFileName(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		// Fallback for invalid URLs
		return strings.ReplaceAll(rawURL, "/", "_")
	}
	// Combine host and path, replacing invalid characters
	name := fmt.Sprintf("%s%s", parsedURL.Host, parsedURL.Path)
	r := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		"&", "_",
		"=", "_",
	)
	return r.Replace(name)
}
