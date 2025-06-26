package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func HandleDownload(input string, output string) {
	data, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	lines := strings.Split(string(data), "\n")

	var urls []string

	// Read URLs
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			urls = append(urls, line)
		}
	}

	wg := &sync.WaitGroup{}

	wg.Add(len(urls))
	// Download Files
	for _, url := range urls {
		go downloadFile(url, output, wg)
	}

	wg.Wait()
}

func downloadFile(url, destDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create destDir if doesn't exist
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to download file: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Skipped %s: HTTP %d\n", url, resp.StatusCode)
	}
	defer resp.Body.Close()

	tokens := strings.Split(url, "/")
	rawName := tokens[len(tokens)-1]
	fileName := strings.Split(rawName, "?")[0] // remove query params
	if fileName == "" {
		fileName = fmt.Sprintf("file-%d", time.Now().UnixNano())
	}
	fmt.Println("filename:", fileName)

	filePath := filepath.Join(destDir, fileName)
	out, err := os.Create(filePath)

	if err != nil {
		log.Fatalf("Failed to create filePath: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
}
