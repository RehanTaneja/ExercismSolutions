package parsinglogfiles

import (
    "regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile("^\\[(TRC|DBG|INF|WRN|ERR|FTL)\\]")
    if len(re.FindStringIndex(text))<=0{
        return false
    }
    return re.FindStringIndex(text)[0]==0
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile("<[~*=-]*>")
    return re.Split(text,-1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
    var cnt int = 0
    for _,x := range lines{
        cnt += len(re.FindAllString(x,-1))
    }
    return cnt
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile("end-of-line+\\d*")
    a := re.Split(text,-1)
    s := ""
    for _,x := range a{
        s = s + x
    }
    return s
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+([A-Za-z0-9]+)`)
    final := []string{}
    for _, line := range lines {
        matches := re.FindStringSubmatch(line)
        if len(matches) > 1 {
            username := matches[1]
            tagged := "[USR] " + username + " " + line
            final = append(final, tagged)
        } else {
            final = append(final, line)
        }
    }
    return final
}

