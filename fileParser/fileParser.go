package fileParser

import (
	"bufio"
	"go-wc/flags"
	"os"
	"strings"
	"unicode/utf8"
)

type FileName string

// ParseFile analyzes a file and returns results for the specified flags
func ParseFile(fileName FileName, enabledFlags []flags.FlagType) ([]flags.Result, error) {
	file, err := os.Open(string(fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read entire file content
	scanner := bufio.NewScanner(file)
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\n") // Add back the newline
	}

	fileContent := content.String()

	results := make(map[flags.FlagType]*flags.Result)

	// Initialize results for enabled flags
	for _, flagType := range enabledFlags {
		results[flagType] = &flags.Result{Type: flagType}
	}

	// Count everything from the full content
	if _, exists := results[flags.LinesFlag]; exists {
		lines := strings.Split(strings.TrimSuffix(fileContent, "\n"), "\n")
		results[flags.LinesFlag].Value = len(lines)
		results[flags.LinesFlag].Label = "lines"
	}

	if _, exists := results[flags.WordsFlag]; exists {
		words := strings.Fields(fileContent)
		results[flags.WordsFlag].Value = len(words)
		results[flags.WordsFlag].Label = "words"
	}

	if _, exists := results[flags.CharsFlag]; exists {
		results[flags.CharsFlag].Value = utf8.RuneCountInString(fileContent)
		results[flags.CharsFlag].Label = "chars"
	}

	if _, exists := results[flags.BytesFlag]; exists {
		results[flags.BytesFlag].Value = len(fileContent)
		results[flags.BytesFlag].Label = "bytes"
	}

	// Convert map to slice
	var resultSlice []flags.Result
	for _, result := range results {
		resultSlice = append(resultSlice, *result)
	}

	return resultSlice, scanner.Err()
}
