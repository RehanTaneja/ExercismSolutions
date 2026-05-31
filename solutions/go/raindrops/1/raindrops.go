package raindrops

import(
    "fmt"
)

func Convert(number int) string {
	answer := ""
    if number%3==0{
        answer = answer + "Pling"
    }
    if number%5==0{
        answer = answer + "Plang"
    }
    if number%7==0{
        answer += "Plong"
    }
    if number%3!=0 && number%5!=0 && number%7!=0{
        return fmt.Sprintf("%v",number)
    }
    return answer
}
