package main

import (
	"errors"
	"flag"
	"os"
)

type inputFile struct {
	filepath  string
	separator string
	valid     bool
}

func main() {

}

func getFileData() (inputFile, error) {

	// Make sure we receive correct amount of arguments
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("filepath argument is required")
	}

	// flag declarations for CLI
	separator := flag.String("separator", "comma", "Column separator")
	valid := flag.Bool("valid", false, "Generate valid JSON data")

	flag.Parse()

	// Store file path value from given argument in the CLI
	filePath := flag.Arg(0)

	// Validate that the argument for the separator flag was either "comma" or "semicolon"
	if !(*separator == "comma" || *separator == "semicolon") {
		return inputFile{}, errors.New("only comma and semicolon separators are allowed")
	}

	// return struct with its required validated data
	return inputFile{filePath, *separator, *valid}, nil

}
