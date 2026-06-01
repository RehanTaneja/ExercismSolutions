package diffsquares

func SquareOfSum(n int) int {
	var i int = 1
    var s int = 0
    for i<=n {
        s+=i
        i++
    }
    return s*s
}

func SumOfSquares(n int) int {
	var i int = 1
    var s int = 0
    for i<=n {
        s+= i*i
        i++
    }
    return s
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
