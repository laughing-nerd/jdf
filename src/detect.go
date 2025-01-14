package src

import (
	"github.com/laughing-nerd/jdf/utils"
)

// According to this function, JSON is something that is enclosed within [] or {}
// But, everuthing enclosed within [] or {} might not always be a JSON.
func DetectJSON(s string) (bool, int, int) {
	var (
		startIndex int = -1
		endIndex   int = -1
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

	braceStack := []rune{} // Stack to keep a track of the braces for json detection

	// JSON end index
	for i, v := range s[startIndex:] {
		if utils.IsSymmetricPair(v) && len(braceStack) > 0 {
			lastEle := braceStack[len(braceStack)-1]
			if lastEle == v {
				braceStack = braceStack[:len(braceStack)-1]
			} else {
				braceStack = append(braceStack, v)
			}
		}

		if utils.IsOpeningPair(v) {
			braceStack = append(braceStack, utils.GetPair(v))
			continue
		}

		if utils.IsClosingPair(v) {
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

	return true, startIndex, endIndex
}
