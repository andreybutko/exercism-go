// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	f := strings.TrimSpace(remark)

	if ContainsUpperLetters(f) && EndsWithQuestion(f) {
		return "Calm down, I know what I'm doing!"
	} else if ContainsUpperLetters(f) {
		return "Whoa, chill out!"
	} else if len(strings.TrimSpace(f)) == 0 {
		return "Fine. Be that way!"
	} else if EndsWithQuestion(f) {
		return "Sure."
	} else {
		return "Whatever."
	}
}

func ContainsUpperLetters(s string) bool {
	return strings.ToUpper(s) == s && ContainsLetters(s)
}

// Checks for letters in string.
func ContainsLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func EndsWithQuestion(s string) bool {
	return (len(s) - 1) == strings.LastIndex(s, "?")
}
