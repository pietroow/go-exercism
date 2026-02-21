package logs

import(
    "unicode/utf8"
)

type Info struct {
    Application string
    unicodePoint string
    character rune
}

var table = map[rune]Info{
    '❗': Info{"recommendation", "U+2757", '❗'},
    '🔍': Info{"search", "U+1F50D", '🔍'},
    '☀': Info{"weather", "U+2600", '☀'},
}

// Application identifies the application emitting the given log.
func Application(log string) string {
    for _, char := range log {
        if value, ok := table[char]; ok {
            return value.Application
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
    for idx, runeValue := range runes {
        if runeValue == oldRune {
            runes[idx] = newRune
        }
    }
    return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    numberOfRunes := utf8.RuneCountInString(log)
    return numberOfRunes <= limit
}
