package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[-=~*]*>`).Split(text, -1)
}
func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	ans := 0
	for _, line := range lines {
		if re.MatchString(line) {
			ans++
		}
	}
	return ans
}
func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	val := []string{}
	for _, line := range lines {
		namesFound := re.FindStringSubmatch(line)
		if namesFound != nil {
			line = fmt.Sprintf("[USR] %s %s", namesFound[1], line)
		}
		val = append(val, line)
	}
	return val
}
