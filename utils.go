package main

import (
	"encoding/json"
	"regexp"
)

func removeANSIColors(s string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(s, "")
}

func containsJSON(s string) bool {
	re := regexp.MustCompile(`\{(?:[^{}"]*|"(?:[^"\\]|\\.)*"|"(?:[^"\\]|\\.)*":(?:[^{}"]*|"(?:[^"\\]|\\.)*"|[^{}"]*)*)\}`)
	match := re.FindStringSubmatch(s)

	if len(match) == 0 {
		return false
	}

	var data interface{}
	err := json.Unmarshal([]byte(match[0]), &data)
	if err != nil {
		return false
	}
	return true
}

func getJSON(s string) string {
	var (
		braces int = 0
		start  int = -1
		end    int = -1
	)
	for index, value := range s {
		// 123 -> {
		// 125 -> }
		// 34 -> "
		if value == 123 {
			// Check for one character forward if start has not been assigned an index
			if start == -1 && s[index+1] == 34 {
				start = index
			}
			braces++
		} else if value == 125 {
			braces--
		}

		if start != -1 && braces == 0 {
			end = index
			break
		}
	}

	if end == -1 {
		end = len(s) - 1
	}
	return s[start : end+1]
}

func getFormattedJSON(s string) (string, error) {
	var data interface{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return "", nil
	}
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

// TODO: Improve the regex in this function
func colorizeJSON(jsonStr, keyColor, valueColor string) string {
	// Regular expressions for keys and values
	keyPattern := `(?:"[^"]*":)`
	valuePattern := `(?::\s*("[^"]*"|\d+|true|false|null))`

	// Apply color to keys
	jsonStr = regexp.MustCompile(keyPattern).ReplaceAllStringFunc(jsonStr, func(match string) string {
		return keyColor + match
	})

	// Apply color to values
	jsonStr = regexp.MustCompile(valuePattern).ReplaceAllStringFunc(jsonStr, func(match string) string {
		return valueColor + match + Colors["reset"]
	})

	return jsonStr
}
