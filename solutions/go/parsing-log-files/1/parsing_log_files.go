package parsinglogfiles
import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"[^"]*(?i:password)[^"]*"`)
    count := 0
    for _, val := range lines {
        if re.MatchString(val) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line(\d)*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, val := range lines {
        match := re.FindStringSubmatch(val)
        if len(match) > 1 {
            lines[i] = fmt.Sprintf("[USR] %s %s", match[1], val)
        }
    }
    return lines
}
