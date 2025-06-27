package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("File compressor CLI started...")

	if len(os.Args) < 2 {
		fmt.Println("Expected at least one subcommand: download, compress, analyze")
		os.Exit(1)
	}

	subCommand := os.Args[1]

	switch subCommand {
	case "download":
		downloadSubCmd := flag.NewFlagSet("download", flag.ExitOnError)
		input := downloadSubCmd.String("input", "", "Path to file with URLs")
		output := downloadSubCmd.String("output", "", "Path to store downloaded files")
		shouldAnalyze := downloadSubCmd.Bool("analyze", false, "Should analyze downloaded items")

		downloadSubCmd.Parse(os.Args[2:])

		HandleDownload(*input, *output, *shouldAnalyze)
	case "compress":
		compressSubCmd := flag.NewFlagSet("compress", flag.ExitOnError)
		input := compressSubCmd.String("input", "", "Path to dir with files")
		output := compressSubCmd.String("output", "", "Path to zipped output directory")

		compressSubCmd.Parse(os.Args[2:])

		HandleCompress(*input, *output)
	case "analyze":
		fmt.Println("Analyze")
	default:
		fmt.Println("Unknown subcommand:", subCommand)
		os.Exit(1)
	}
}
