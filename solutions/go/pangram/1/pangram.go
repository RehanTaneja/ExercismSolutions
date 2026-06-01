package pangram

import(
    "strings"
    "unicode"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
    a := make([]int,26)
    for _,x := range input{
        if unicode.IsLetter(x){
            a[x-'a']++
        }
    }
    for _,x := range a{
        if x==0{
            return false
        }
    }
    return true
}
