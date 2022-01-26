package logs
import (
    "unicode/utf8"
    //"fmt"
    //"strconv"
    )

// Application identifies the application emitting the given log.
func Application(log string) string {
    app := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001f50d': "search",
		'\u2600':     "weather",
	}
	for _,i:= range(log) {
           v,e := app[i]
        if e {
            return v
        }
    }
return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    s := []rune(log)
    for i,v := range(s) {
        if v == oldRune {
            s[i] = newRune
        }
    }
return string(s)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	l := utf8.RuneCountInString(log)
    if l <= limit {
        return true
    } else {
    return false
    }
}
