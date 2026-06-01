package romannumerals

import(
    "errors"
)

func ToRomanNumeral(input int) (string, error) {
    if (input<=0 || input>=4000){
        return "",errors.New("Invalid Int")
    }
	var a string = ""
    for input>0{
        if input >=1000{
            a += "M"
            input -= 1000
        }else if input >= 900{
            a += "CM"
            input -= 900
        }else if input >= 500{
            a += "D"
            input -= 500
        }else if input >= 400{
            a += "CD"
            input -= 400
        }else if input >= 100{
            a += "C"
            input -= 100
        }else if input >= 90{
            a += "XC"
            input -= 90
        }else if input >= 50{
            a += "L"
            input -= 50
        }else if input >= 40{
            a += "XL"
            input -= 40
        }else if input >= 10{
            a += "X"
            input -= 10
        }else if input >= 9{
            a += "IX"
            input -= 9
        }else if input >= 5{
            a += "V"
            input -= 5
        }else if input >= 4{
            a += "IV"
            input -= 4
        }else if input >= 1{
            a += "I"
            input -= 1
        }
    }
    return a,nil
}
