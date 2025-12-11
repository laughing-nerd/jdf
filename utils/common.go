package utils

const (
	CURLY_BRACES_OPEN   byte = '{'
	CURLY_BRACES_CLOSE  byte = '}'
	SQUARE_BRACES_OPEN  byte = '['
	SQUARE_BRACES_CLOSE byte = ']'
	DOUBLE_QUOTE        byte = '"'
	COLON               byte = ':'
	COMMA               byte = ','
	BACKSLASH           byte = '\\'
)

func GetPair(val byte) byte {
	switch val {
	case CURLY_BRACES_OPEN:
		return CURLY_BRACES_CLOSE
	case SQUARE_BRACES_OPEN:
		return SQUARE_BRACES_CLOSE
	case DOUBLE_QUOTE:
		return DOUBLE_QUOTE
	default:
		return 0
	}
}

func IsOpeningPair(val byte) bool {
	switch val {
	case CURLY_BRACES_OPEN, SQUARE_BRACES_OPEN:
		return true
	default:
		return false
	}
}

func IsClosingPair(val byte) bool {
	switch val {
	case CURLY_BRACES_CLOSE, SQUARE_BRACES_CLOSE:
		return true
	default:
		return false
	}
}

// IsEscaped returns true if the character at position i is escaped.
// It counts consecutive backslashes before position i.
// An odd number of backslashes means the character is escaped.
func IsEscaped(s string, i int) bool {
	if i <= 0 {
		return false
	}
	backslashCount := 0
	for j := i - 1; j >= 0 && s[j] == BACKSLASH; j-- {
		backslashCount++
	}
	return backslashCount%2 == 1
}
