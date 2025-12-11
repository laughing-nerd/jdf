package utils

import (
	"flag"
	"strings"
)

var (
	FlagSeparator    string
	FlagIndent       int64
	FlagWebMode      bool
	FlagPort         int64
	FlagNoEscape     bool
	FlagIndentGuides bool
)

func RegisterFlags() {

	flag.StringVar(&FlagSeparator, "separator", "=", "Sets the separator")
	flag.StringVar(&FlagSeparator, "s", "=", "Sets the separator (shorthand)")

	flag.Int64Var(&FlagIndent, "indent", 2, "Sets the indentation")
	flag.Int64Var(&FlagIndent, "i", 2, "Sets the indentation (shorthand)")

	flag.Int64Var(&FlagPort, "port", 6969, "Sets the port")
	flag.Int64Var(&FlagPort, "p", 6969, "Sets the port (shorthand)")

	flag.BoolVar(&FlagWebMode, "web", false, "Formats logs and outputs in HTML. Logs are updated in real-time")

	flag.BoolVar(&FlagNoEscape, "no-escape", false, "Do not escape special characters")
	flag.BoolVar(&FlagNoEscape, "ne", false, "Do not escape special characters (shorthand)")

	flag.BoolVar(&FlagIndentGuides, "guides", false, "Show indent guide lines")
	flag.BoolVar(&FlagIndentGuides, "g", false, "Show indent guide lines (shorthand)")

	flag.Parse()
}

// Box-drawing characters for tree-style indent guides
const (
	boxVertical = "\u2502" // │
	boxSpace    = " "
)

// GetIndentString returns the indentation string for the given level
// If FlagIndentGuides is true, it includes vertical guide lines like tree command
func GetIndentString(level int) string {
	if level <= 0 {
		return ""
	}

	indentWidth := int(FlagIndent)
	if indentWidth < 1 {
		indentWidth = 2
	}

	// simple spaces
	if !FlagIndentGuides {
		totalSpaces := level * indentWidth
		return strings.Repeat(" ", totalSpaces)
	}

	// With indent guides: "│ " repeated for each level (tree-style)
	var result string
	// Using dark gray (256-color: 239) for subtle guides
	guideColor := "\033[38;5;239m"
	resetColor := "\033[0m"

	for range level {
		result += guideColor + boxVertical + resetColor
		// add spaces after the vertical line (indentWidth - 1 spaces)
		for range indentWidth - 1 {
			result += boxSpace
		}
	}
	return result
}
