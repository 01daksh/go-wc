package main

import (
	"fmt"
	"go-wc/fileParser"
	"go-wc/wcFlag"
)

func GetFileStats() {
	options := wcFlag.ParseFlags()

	for _, file := range options.Files {
		results, err := fileParser.ParseFile(file, options.EnabledFlags)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", file, err)
			continue
		}

		for _, result := range results {
			fmt.Printf("%d %s ", result.Value, result.Label)
		}
		fmt.Printf("%s\n", file)
	}
}

func main() {
	GetFileStats()
}
