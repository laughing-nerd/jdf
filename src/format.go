package src

import (
	"strings"

	"github.com/laughing-nerd/jdf/utils"
)

func FormatJSON(s string) string {
	var (
		indentCount int64  = utils.Indent
		indent      string = ""
		temp        string = ""
		result      string = ""
		quoteStatus int    = -1 // holds the double quote status. -1 means the double quote is not open
	)

	for _, v := range s {

		if v == utils.DOUBLE_QUOTE {
			quoteStatus *= -1
		}

		// indent only if the opening pair is not a part of temp string
		if utils.IsOpeningPair(v) && quoteStatus == -1 {
			result += string(v) + "\n"
			indent = updateIndent(indent, indentCount)
			result += indent
			continue
		}

		// add color to the key part of the key-value pair if colon is separating key and value
		// do nothing if the colon is a part of temp string
		if v == utils.COLON && quoteStatus == -1 {
			result += utils.Colorize(temp, "blue") + ": "
			temp = ""
			continue
		}

		// indent on comma only if the comma is separating two key-value pairs
		if v == utils.COMMA && quoteStatus == -1 {
			result += utils.Colorize(temp, getColor(temp)) + string(v) + "\n" + indent
			temp = ""
			continue
		}

		// indent only if the closing pair is not a part of temp string
		if utils.IsClosingPair(v) && quoteStatus == -1 {
			indent = updateIndent(indent, -indentCount)
			result += utils.Colorize(temp, getColor(temp)) + "\n" + indent + string(v)
			temp = ""
			continue
		}

		temp += string(v)
	}

	return result
}

// helper func ...
func updateIndent(indent string, count int64) string {
	if count > 0 {
		for i := int64(0); i < count; i++ {
			indent += " "
		}
	} else {
		if len(indent) >= int(count*-1) {
			indent = indent[:len(indent)-int(count*-1)]
		}
	}
	return indent
}

func getColor(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return "orange"
	} else {
		return "green"
	}
}
