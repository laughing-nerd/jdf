package utils

const (
	CURLY_BRACES_OPEN   rune = '{'
	CURLY_BRACES_CLOSE  rune = '}'
	SQUARE_BRACES_OPEN  rune = '['
	SQUARE_BRACES_CLOSE rune = ']'
	DOUBLE_QUOTE        rune = '"'
	COLON               rune = ':'
	COMMA               rune = ','
	BACKSLASH           rune = '\\'
)

func GetPair(val rune) rune {
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

func IsOpeningPair(val rune) bool {
	switch val {
	case CURLY_BRACES_OPEN, SQUARE_BRACES_OPEN:
		return true
	default:
		return false
	}
}

func IsClosingPair(val rune) bool {
	switch val {
	case CURLY_BRACES_CLOSE, SQUARE_BRACES_CLOSE:
		return true
	default:
		return false
	}
}
