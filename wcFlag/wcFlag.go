package wcFlag

import (
	"flag"
	"go-wc/fileParser"
	"go-wc/flags"
)

type FlagOptions struct {
	EnabledFlags []flags.FlagType
	Files        []fileParser.FileName
}

func ParseFlags() FlagOptions {
	options := FlagOptions{}

	// Create variables to hold flag values
	var linesFlag, wordsFlag, charsFlag, bytesFlag bool

	// Register flags with their variables
	flag.BoolVar(&linesFlag, "l", false, "print the number of lines")
	flag.BoolVar(&linesFlag, "lines", false, "print the number of lines")
	flag.BoolVar(&wordsFlag, "w", false, "print the number of words")
	flag.BoolVar(&wordsFlag, "words", false, "print the number of words")
	flag.BoolVar(&charsFlag, "m", false, "print the number of characters")
	flag.BoolVar(&charsFlag, "chars", false, "print the number of characters")
	flag.BoolVar(&bytesFlag, "c", false, "print the number of bytes")
	flag.BoolVar(&bytesFlag, "bytes", false, "print the number of bytes")

	flag.Parse()

	// Collect enabled flags based on the flag variables
	if linesFlag {
		options.EnabledFlags = append(options.EnabledFlags, flags.LinesFlag)
	}
	if wordsFlag {
		options.EnabledFlags = append(options.EnabledFlags, flags.WordsFlag)
	}
	if charsFlag {
		options.EnabledFlags = append(options.EnabledFlags, flags.CharsFlag)
	}
	if bytesFlag {
		options.EnabledFlags = append(options.EnabledFlags, flags.BytesFlag)
	}

	// Handle files
	options.Files = make([]fileParser.FileName, len(flag.Args()))
	for i, arg := range flag.Args() {
		options.Files[i] = fileParser.FileName(arg)
	}

	return options
}
