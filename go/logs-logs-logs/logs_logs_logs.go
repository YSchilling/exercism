package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range strings.Split(log, "") {
		if v == "❗" {
			return "recommendation"
		} else if v == "🔍" {
			return "search"
		} else if v == "☀" {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	splittedLog := strings.Split(log, "")
	for i, v := range splittedLog {
		if v == string(oldRune) {
			splittedLog[i] = string(newRune)
		}
	}
	return strings.Join(splittedLog, "")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
