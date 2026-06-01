package armstrongnumbers

import "math"

func IsNumber(n int) bool {
	if n<10 {
        return true
    }
    var num int = n
    var digits int = 0
    for num > 0 {
        num = num/10
        digits++
    }
    var number int = n
    var sum int = 0
    for n > 0 {
        var d int = n%10
        n = n/10
        sum += int(math.Pow(float64(d),float64(digits)))
    }
    return sum == number
}
