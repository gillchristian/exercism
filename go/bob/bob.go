package bob // package name must match the package name in bob_test.go

import (
	"regexp"
	"strings"
)

const testVersion = 2 // same as targetTestVersion

func Hey(message string) string {
	message = strings.TrimSpace(message)

	if len(message) == 0 {
		return "Fine. Be that way!"
	}

	if isNothig(message) && !isQuestion(message) {
		return "Whatever."
	}

	if isYelling(message) && !isNothig(message) {
		return "Whoa, chill out!"
	}

	if isQuestion(message) {
		return "Sure."
	}

	return "Whatever."
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

func isNothig(s string) bool {
	r, _ := regexp.Compile("([A-Z]|[a-z])")
	return !r.MatchString(s)
}

func isYelling(s string) bool {
	return strings.ToUpper(s) == s
}
