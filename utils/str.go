package utils

import "regexp"

func RemoveWhitespace(input string) string {
	// Create a regular expression pattern to match whitespace
	pattern := "\\s+"
	regex := regexp.MustCompile(pattern)

	// Replace all matches with an empty string
	result := regex.ReplaceAllString(input, "")

	return result
}
