package flags

type FlagType string

const (
	LinesFlag FlagType = "lines"
	WordsFlag FlagType = "words"
	CharsFlag FlagType = "chars"
	BytesFlag FlagType = "bytes"
)

// FlagDefinition represents a single flag configuration
type FlagDefinition struct {
	Type        FlagType
	ShortName   string
	LongName    string
	Description string
	Enabled     bool
}

// FlagRegistry holds all available flags
type FlagRegistry map[FlagType]FlagDefinition

// Result holds the computed value for a flag
type Result struct {
	Type  FlagType
	Value int
	Label string
}
