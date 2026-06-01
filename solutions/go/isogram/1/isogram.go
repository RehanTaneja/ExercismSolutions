package isogram

import(
    "strings"
    "unicode"
)

func IsIsogram(word string) bool {
    word = strings.ToUpper(word)
	mymap := make(map[rune]int)
    for _,x := range word{
        if mymap[x]>=1 && unicode.IsLetter(x){
            return false
        }
        mymap[x] = 1
    }
    return true
}
