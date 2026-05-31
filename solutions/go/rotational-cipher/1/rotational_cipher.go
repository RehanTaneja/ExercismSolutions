package rotationalcipher

import(
    "unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	s := ""
    for _,x := range plain{
        if unicode.IsLetter(x){
            var d rune;
            if unicode.IsUpper(x){
                d = (x-'A'+rune(shiftKey))%26 + 'A'
            }else{
                d = (x-'a'+rune(shiftKey))%26 + 'a'
            }
            s += string(d)
        }else{
            s += string(x)
        }
    }
    return s
}
