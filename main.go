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

		downloadSubCmd.Parse(os.Args[2:])

		HandleDownload(*input, *output)
	case "compress":
		fmt.Println("Compress")
	case "analyze":
		fmt.Println("Analyze")
	default:
		fmt.Println("Unknown subcommand:", subCommand)
		os.Exit(1)
	}
}
