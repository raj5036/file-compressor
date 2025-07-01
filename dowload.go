package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func HandleDownload(input, output string, shouldAnalyze bool) {
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

	const maxConcurrentDownloads = 5
	semaphore := make(chan struct{}, maxConcurrentDownloads)
	wg := &sync.WaitGroup{}

	// Download Files
	for _, url := range urls {
		semaphore <- struct{}{}
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			downloadFile(u, output)

			<-semaphore
		}(url)

	}

	wg.Wait()

	if shouldAnalyze {
		totalFileCount, totalFileSize, fileExtCounts, analyze_err := analyzeDownloadDirectory(output)
		if analyze_err != nil {
			log.Fatalf("Error analyzing downloaded files: %v\n", analyze_err)
		}

		fmt.Println("\n💻Analyzed output:")
		fmt.Printf("✅ Total files: %v\n", totalFileCount)
		fmt.Printf("📦 Total size: %v mb\n", math.Floor(totalFileSize*100)/100)

		fmt.Println("🗂️ File types:")
		for extension, count := range fileExtCounts {
			fmt.Printf("%s: %d\n", extension, count)
		}
	}
}

func downloadFile(url, destDir string) {

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

func analyzeDownloadDirectory(destDir string) (int, float64, map[string]int, error) {
	var totalFileCount int = 0
	var totalFileSize float64 = 0 // In mb
	fileExtCounts := make(map[string]int)

	err := filepath.Walk(destDir, func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() { // TODO: Handle nested directories
			return nil
		}

		totalFileCount++
		totalFileSize += float64(fileInfo.Size())

		fileExtension := filepath.Ext(fileInfo.Name())
		if fileExtCounts[fileExtension] == 0 {
			fileExtCounts[fileExtension] = 1
		} else {
			fileExtCounts[fileExtension]++
		}

		return err
	})

	return totalFileCount, totalFileSize / 1024.0 / 1024.0, fileExtCounts, err
}
