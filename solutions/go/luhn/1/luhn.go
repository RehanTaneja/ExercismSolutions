package luhn

import (
    "strings"
    "unicode"
)

func Valid(id string) bool {
    id = strings.ReplaceAll(id," ","")
    if len(id)<=1{return false}
	runes := []rune(id)
    num := 3
    sum := 0
    for i := len(runes)-1;i>=0;i--{
        if (unicode.IsDigit(runes[i])){
            d := int(runes[i]-'0')
        	if num%2==0{
            	d = d*2
            	if d>9{d-=9}
        	}
            num++
        	sum += d
        }else{
            return false
        }
    }
    return sum%10 == 0
    
}
