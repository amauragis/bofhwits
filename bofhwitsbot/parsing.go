package bofhwitsbot

import (
	// "fmt"

	"strconv"
	"strings"
)

// format the string (probably a tweet request) for a single line
func formatTextOneLine(s string) string {
	formatStr := strings.TrimSpace(s)
	return strings.Replace(formatStr, "\n", " | ", -1)
}

// format the string (log request for site) preserving line endings
func formatTextMultiLine(s string) string {
	return s
}

func separateUsername(s string) (user string, msg string) {

	// search for a < and > pair and take everything between them
	if strings.ContainsRune(s, '<') {
		userStart := strings.IndexRune(s, '<') + 1
		userEnd := strings.IndexRune(s, '>')

		user = strings.TrimSpace(s[userStart:userEnd])
		msg = strings.TrimSpace(s[userEnd+1:])

	} else {
		// handle single ended delimiters
		delims := []rune{'>', ':'}
		var matchRune rune
		for _, elem := range delims {

			if strings.ContainsRune(s, elem) {
				matchRune = elem
			}
		}

		quotedRuneStr := strconv.QuoteRuneToASCII(matchRune)
		runeStr := quotedRuneStr[1 : len(quotedRuneStr)-1]
		splitstr := strings.Split(s, runeStr)

		user = strings.TrimSpace(splitstr[0])
		msg = strings.TrimSpace(splitstr[1])
	}

	return
}

func testSubmissionValidity(s string) bool {
	delims := []rune{'>', ':'}
	exists := false
	for _, elem := range delims {
		if strings.ContainsRune(s, elem) {
			exists = true
		}
	}
	return exists
}
