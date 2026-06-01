package logs

import (
    "fmt"
    "unicode/utf8"
    )

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _,v := range log{
        if fmt.Sprintf("%U",v) == "U+2757"{
            return "recommendation"
        }else if fmt.Sprintf("%U",v) == "U+1F50D"{
            return "search"
        }else if fmt.Sprintf("%U",v) == "U+2600"{
            return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    mylog := []rune(log)
	for i,v := range mylog{
        if fmt.Sprintf("%U",v)==fmt.Sprintf("%U",oldRune){
            mylog[i] = newRune
        }
    }
    log = string(mylog)
    return log
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log)<=limit
}
