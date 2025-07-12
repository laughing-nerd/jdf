package term

import (
	"strings"
)

var separatorStr string

func InitSeparator(sep string) string {
	width, err := getTerminalWidth()
	if err != nil {
		width = 80
	}
	return strings.Repeat(sep, width)
}
