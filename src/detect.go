package src

import (
	"github.com/laughing-nerd/jdf/utils"
)

// According to this function, JSON is something that is enclosed within [] or {}
// But, everuthing enclosed within [] or {} might not always be a JSON.
func DetectJSON(s string) (bool, int, int) {
	var (
		startIndex int    = -1
		endIndex   int    = -1
		braceStack []byte = []byte{} // Stack to keep a track of the braces for json detection
		quoteStatus bool = false // holds the double quote status. -1 means the double quote is not open
	)

	// JSON start index
	for i := range s {

		// if the character is [ and it is not followed by {, then it is not a JSON
		if s[i] == utils.SQUARE_BRACES_OPEN && i+1 < len(s) && s[i+1] != utils.CURLY_BRACES_OPEN {
			continue
		}

		if !utils.IsOpeningPair(s[i]) {
			continue
		}

		startIndex = i
		break
	}

	// If the JSON start index is not found or the startIndex is the last index of the string, then it is not a JSON
	if startIndex == -1 || startIndex == len(s)-1 {
		return false, startIndex, endIndex
	}

	// JSON end index
	// for i, v := range s[startIndex:] {
	for i := startIndex; i < len(s); i++ {

		// if the character is " and is not preceded by \, then toggle the quote status
		if s[i] == utils.DOUBLE_QUOTE && s[i-1] != utils.BACKSLASH {
			quoteStatus = !quoteStatus
		}

		if !quoteStatus {

			// consider the opening pair only if the double quote status is not open
			if utils.IsOpeningPair(s[i]) {
				braceStack = append(braceStack, utils.GetPair(s[i]))
				continue
			}

			// consider the closing pair only if the double quote status is not open
			if utils.IsClosingPair(s[i]) {
				lastEle := braceStack[len(braceStack)-1]
				if lastEle != s[i] {
					return false, startIndex, endIndex
				}
				braceStack = braceStack[:len(braceStack)-1]
			}
		}

		// if brace stack is empty then that means there are no more braces to consider
		if len(braceStack) == 0 {
			endIndex = i
			break
		}
	}

	// If the brace stack is not empty or double quote is opened, then it is not a JSON
	if len(braceStack) > 0 || quoteStatus {
		return false, startIndex, endIndex
	}

	return true, startIndex, endIndex
}
