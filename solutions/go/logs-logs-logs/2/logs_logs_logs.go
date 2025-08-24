package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, val := range log {
        switch val{
            case '\u2757': return "recommendation"
            case '\U0001F50D': return "search"
            case '\u2600': return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var output []rune
    for _, val := range log {
        if val == oldRune {
            output = append(output, newRune)
        } else {
            output = append(output, val)
        }
    }
    return string(output)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
