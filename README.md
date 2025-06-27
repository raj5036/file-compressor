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

## 📌 Requirements
- Go 1.18+
- Internet access for downloads
- Write permission to output folder