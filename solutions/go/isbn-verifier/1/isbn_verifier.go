package isbn

func IsValidISBN(isbn string) bool {
	var num int = 10
    var s int = 0
    for _,x := range isbn{
        if x=='-'{
            continue
        }else if x-'0'>=0 && x-'0'<=9{
            s += num*int(x-'0')
            num -= 1
        }else if x=='X'{
            if num!=1{
                return false
            }
            s += 10
            num-=1
        }else{
            return false
        }
    }
    return (s%11==0) && num==0
}
