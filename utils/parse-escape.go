package utils

import "strconv"

// Map of escape sequences to their corresponding characters
var escapeSequences = map[string]string{
	`\n`: "\n", // Newline
	`\r`: "\r", // Carriage Return
	`\t`: "\t", // Tab
	// '\"': `"`,  // Double Quote
	`\\`: `\`, // Backslash
	// \/': "/",  // Forward Slash
	`\'`: `'`,   // Single Quote
	`\b`: "\b", // Backspace
	`\f`: "\f", // Form Feed
	`\v`: "\v", // Vertical Tab
}

func ParseEscape(s string, i int) (string, int) {
	if i+1 < len(s) && s[i+1] == 'u' {
		// Ensure there are at least 4 more characters after \u
		if i+5 < len(s) {
			hexPart := s[i+2 : i+6] // Extract only the hex digits
			codePoint, err := strconv.ParseInt(hexPart, 16, 32)
			if err != nil {
				return s[i : i+6], 5 // Return original if parsing fails
			}
			return string(rune(codePoint)), 5
		}
	}

	// Handle normal escape sequences
	if val, ok := escapeSequences[s[i:i+2]]; ok {
		return val, 1
	}
	
	return string(s[i]), 0
}

