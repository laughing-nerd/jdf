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
