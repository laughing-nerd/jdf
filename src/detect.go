package src

import (
	"github.com/laughing-nerd/jdf/utils"
)

// According to this function, JSON is something that is enclosed within [] or {}
// But, everuthing enclosed within [] or {} might not always be a JSON.
func DetectJSON(s string) (bool, int, int) {
	var (
		startIndex  int    = -1
		endIndex    int    = -1
		braceStack  []rune = []rune{} // Stack to keep a track of the braces for json detection
		quoteStatus int    = -1       // holds the double quote status. -1 means the double quote is not open
	)

	// JSON start index
	for i, v := range s {

		// if the character is [ and it is not followed by {, then it is not a JSON
		if v == utils.SQUARE_BRACES_OPEN && i+1 < len(s) && rune(s[i+1]) != utils.CURLY_BRACES_OPEN {
			continue
		}

		if !utils.IsOpeningPair(v) {
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
	for i, v := range s[startIndex:] {

		// if the character is " and is not preceded by \, then toggle the quote status
		if v == utils.DOUBLE_QUOTE && rune(s[i-1]) != utils.BACKSLASH {
			quoteStatus *= -1
		}

		// consider the opening pair only if the double quote status is not open
		if utils.IsOpeningPair(v) && quoteStatus == -1 {
			braceStack = append(braceStack, utils.GetPair(v))
			continue
		}

		// consider the closing pair only if the double quote status is not open
		if utils.IsClosingPair(v) && quoteStatus == -1 {
			lastEle := braceStack[len(braceStack)-1]
			if lastEle != v {
				return false, startIndex, endIndex
			}
			braceStack = braceStack[:len(braceStack)-1]
		}

		// if brace stack is empty then that means there are no more braces to consider
		if len(braceStack) == 0 {
			endIndex = i + startIndex
			break
		}
	}

	if quoteStatus == 1 {
		return false, startIndex, endIndex
	}

	return true, startIndex, endIndex
}
