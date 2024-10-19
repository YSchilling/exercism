package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(~|\*|=|-)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	counter := 0
	re := regexp.MustCompile(`(?i)"[[:ascii:]]*password[[:ascii:]]*"`)

	for _, v := range lines {
		if re.MatchString(v) {
			counter += 1
		}
	}

	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +([a-zA-Z0-9]+)`)
	var newLines []string

	for _, v := range lines {
		matches := re.FindStringSubmatch(v)

		if len(matches) > 0 {
			username := matches[1]
			formattedLine := fmt.Sprintf("[USR] %s %s", username, v)

			newLines = append(newLines, formattedLine)
		} else {
			newLines = append(newLines, v)
		}
	}

	return newLines
}
