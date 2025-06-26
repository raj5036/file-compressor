package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func HandleCompress(sourceDir, destinationFile string) {
	err := compressDirectory(sourceDir, destinationFile)
	if err != nil {
		log.Fatalf("Error compressing: %v", err)
	}

	log.Printf("Directory compressed successfully to  %v", destinationFile)
}

func compressDirectory(sourceDir, destinationFile string) error {
	// Create the output .tar.gz file
	outputFile, err := os.Create(destinationFile)
	if err != nil {
		log.Printf("Error in creating destination file: %v\n", err)
		return err
	}
	defer outputFile.Close()

	gzWriter := gzip.NewWriter(outputFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	err = filepath.Walk(sourceDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Get relative path for tar header
		relPath := strings.TrimPrefix(path, sourceDir)
		relPath = strings.TrimPrefix(relPath, string(filepath.Separator))

		// Create tar header
		header, err := tar.FileInfoHeader(info, relPath)
		if err != nil {
			return err
		}
		header.Name = relPath

		// Write header
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// Open file to copy content
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Copy file data into tar writer
		_, err = io.Copy(tarWriter, file)
		return err
	})

	return err
}

// func compressFile(sourcePath, destinationPath string) error {
// 	inputFile, err := os.Open(sourcePath)
// 	if err != nil {
// 		log.Fatalf("Error in opening Source directory: %v\n", err)
// 	}
// 	defer inputFile.Close()

// 	outputFile, err := os.Create(destinationPath)
// 	if err != nil {
// 		log.Fatalf("Error in creating destination path: %v", err)
// 	}
// 	defer outputFile.Close()

// 	gzWriter := gzip.NewWriter(outputFile)
// 	defer gzWriter.Close()

// 	// gzWriter.Name = filePath.Base(sourcePath)

// 	_, err = io.Copy(gzWriter, inputFile)
// 	return err
// }
