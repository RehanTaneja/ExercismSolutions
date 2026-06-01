package grains

import(
    "errors"
)

func Square(number int) (uint64, error) {
	if number<=0 || number>64{
        return 0,errors.New("Invalid square number.")
    }
    var a uint64 = 0
    i := 1
    for i <=number{
        if a==0{
            a = 1
        }else{
            a = a * 2
        }
        i++
    }
    return a, nil
}

func Total() uint64 {
	var s uint64 = 0
    i :=1
    for i<=64{
        v,_ := Square(i)
        s += v
        i++
    }
    return s
}
