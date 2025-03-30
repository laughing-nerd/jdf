package src

import (
	"bytes"
	"strings"

	"github.com/laughing-nerd/jdf/utils"
)

func FormatJSON(s string) string {
	var (
		buffer      bytes.Buffer
		indent      int  = 0
		inArray     int  = 0    // tracks array depth. Necessary to determine whether the next thing is a key or not
		keyStatus   bool = true // keyStatus true means the next thing is a key
		quoteStatus bool        // quoteStatus true means (") is open
	)

	// Preallocate memory for buffer to avoid frequent resizes
	buffer.Grow(len(s) + 69)

	for i := 0; i < len(s); i++ {
		char := s[i]

		// If the user has opted for no escape, then this check will not be performed
		if !utils.FlagNoEscape && char == utils.BACKSLASH && quoteStatus {
			escapeParsed, skip := utils.ParseEscape(s, i)
			buffer.WriteString(escapeParsed)
			i += skip
			continue
		}

		// if the character is " and is not preceded by \ or is the first character, then toggle the quote status since this is a part of json format
		if char == utils.DOUBLE_QUOTE && (i == 0 || s[i-1] != utils.BACKSLASH) {
			quoteStatus = !quoteStatus

			// There's probably a better way to do this. But as of now, this is the best way I can think of
			if quoteStatus {
				if keyStatus {
					buffer.WriteString(utils.Colorize("", "blue"))
				} else {
					buffer.WriteString(utils.Colorize("", "orange"))
				}
				buffer.WriteByte(char)
			} else {
				buffer.WriteByte(char)
				buffer.WriteString(utils.Colorize("", "reset"))
			}
			continue
		}

		switch char {
		case utils.CURLY_BRACES_OPEN, utils.SQUARE_BRACES_OPEN:
			buffer.WriteByte(char)
			if !quoteStatus {
				indent++
				buffer.WriteByte('\n')
				buffer.WriteString(strings.Repeat("  ", indent))

				// Increment array depth if it's a [. But when there's a {, it's a key. So set keyStatus to true
				if char == utils.SQUARE_BRACES_OPEN {
					inArray++
				} else {
					inArray = 0
					keyStatus = true
				}
			}

		case utils.CURLY_BRACES_CLOSE, utils.SQUARE_BRACES_CLOSE:
			if !quoteStatus {
				indent--
				buffer.WriteByte('\n')
				buffer.WriteString(strings.Repeat("  ", indent))

				// Only decrement inArray if inArray has a certain depth. Else no point in decrementing
				if char == utils.SQUARE_BRACES_CLOSE && inArray > 0 {
					inArray--
				}
			}
			buffer.WriteByte(char)

		case utils.COMMA:
			buffer.WriteByte(char)
			if !quoteStatus {

				// If the comma is not a part of an array, then it's a part of the original json string and a new key will begin from next line
				if inArray == 0 {
					keyStatus = true
				}
				buffer.WriteByte('\n')
				buffer.WriteString(strings.Repeat("  ", indent))
			}

		case utils.COLON:
			buffer.WriteByte(char)
			if !quoteStatus {
				keyStatus = false
				buffer.WriteByte(' ')
			}

		default:
			// Check for boolean
			if !quoteStatus {
				if char == 't' || char == 'T' || char == 'f' || char == 'F' {
					buffer.WriteString(utils.Colorize("", "green"))
					if char == 't' || char == 'T' {
						buffer.Write([]byte(s[i : i+4]))
						i += 3
					} else {
						buffer.Write([]byte(s[i : i+5]))
						i += 4
					}
					buffer.WriteString(utils.Colorize("", "reset"))
					continue
				}

				// Check for null
				if char == 'n' || char == 'N' {
					buffer.WriteString(utils.Colorize("", "purple"))
					buffer.Write([]byte(s[i : i+4]))
					buffer.WriteString(utils.Colorize("", "reset"))
					i += 3
					continue
				}

				// Numbers will be white for now. Might change later
			}

			buffer.WriteByte(char)
		}
	}

	return buffer.String()
}
