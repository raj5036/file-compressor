# 🧰 File Compressor CLI Tool (in Go)

A fast and lightweight command-line tool written in **Go** that allows you to:

✅ Download multiple files concurrently from a list of URLs  
🗜️ Compress all downloaded files into a single `.tar.gz` archive  
📊 Analyze downloaded or compressed files to gather insights like file types, sizes, and counts  

---

## 📦 Features

- Concurrent file downloads using goroutines
- Safe error handling and retry mechanisms
- Directory compression to `.tar.gz` using `tar` and `gzip`
- Directory and archive analysis for file type breakdown, size stats, etc.
- Modular structure — easy to extend

---

## ⌨️ Commands
1. Download and Analyze file: `go run . download -input="urls.txt" -output=downloads/ -analyze=true` (Sample "urls.txt" can be found in the codebase)
2. Compress downloaded files: `go run . compress -input=downloads/ -output=compressed.tar.gz`

---

## 📌 Requirements
- Go 1.18+
- Internet access for downloads
- Write permission to output folder