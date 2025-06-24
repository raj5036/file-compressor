package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func HandleDownload(input string, output string) {
	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
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

	// Download Files
	for _, url := range urls {
		downloadFile(url, output)
	}
}

func downloadFile(url string, destDir string) {
	// Create destDir if doesn't exist
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	tokens := strings.Split(url, "/")
	rawName := tokens[len(tokens)-1]
	fileName := strings.Split(rawName, "?")[0] // remove query params
	if fileName == "" {
		fileName = fmt.Sprintf("file-%d", time.Now().UnixNano())
	}
	fmt.Println("filename", fileName)

	filePath := filepath.Join(destDir, fileName)
	fmt.Println("filePath", filePath)
	out, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
