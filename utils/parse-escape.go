package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// Map of escape sequences to their corresponding characters
var escapeSequences = map[string]string{
	`\n`: "\n", // Newline
	`\r`: "\r", // Carriage Return
	`\t`: "\t", // Tab
	`\"`: `"`,  // Double Quote
	`\\`: `\`,  // Backslash
	`\/`: "/",  // Forward Slash
	`\'`: `'`,  // Single Quote
	`\b`: "\b", // Backspace
	`\f`: "\f", // Form Feed
	`\v`: "\v", // Vertical Tab
}

func ParseEscape(s *string) {
	// Regex to match \uXXXX (Unicode escape sequence) and common escape sequences
	re := regexp.MustCompile(`\\u[0-9A-Fa-f]{4}|\\[nrt"\\bvf]`)

	// Replace all matches using ReplaceAllStringFunc
	*s = re.ReplaceAllStringFunc(*s, func(m string) string {
		if strings.HasPrefix(m, `\u`) {
			// Handle \uXXXX (Unicode escape sequence)
			codePoint, err := strconv.ParseInt(m[2:], 16, 32)
			if err != nil {
				return m // Return original string if error occurs
			}
			return string(rune(codePoint))
		}

		// Return the corresponding value for known escape sequences
		if val, ok := escapeSequences[m]; ok {
			return val
		} else {
			return m // Return the original string if not recognized
		}
	})
}
