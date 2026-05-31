package reverse

func Reverse(input string) string {
	var a string = ""
    for _,x := range input{
        a = string(x) + a
    }
    return a
}
