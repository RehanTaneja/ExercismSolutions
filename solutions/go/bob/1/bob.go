// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import(
    "unicode"
    "strings"
)

func IsUppercase(s string) bool {
    hasLetters := false
    for _, r := range s {
        if unicode.IsLetter(r) {
            hasLetters = true
            if !unicode.IsUpper(r) {
                return false
            }
        }
    }
    return hasLetters
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
    trimmed := strings.TrimSpace(remark)
	if trimmed==""{
        return "Fine. Be that way!"
    }else if IsUppercase(trimmed)&& strings.HasSuffix(trimmed,"?"){
        return "Calm down, I know what I'm doing!";
    }else if IsUppercase(trimmed){
        return "Whoa, chill out!"
    }else if strings.HasSuffix(trimmed,"?"){
        return "Sure."
    }
    return "Whatever."
}
