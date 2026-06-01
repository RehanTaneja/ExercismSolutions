package microblog

func Truncate(phrase string) string {
    runes := []rune(phrase)
	if len(runes)<=5 {
        return string(runes)
    }
    return string(runes[:5])
}
