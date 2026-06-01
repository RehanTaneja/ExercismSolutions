package darts

func Score(x, y float64) int {
	var a float64 = (x*x) + (y*y)
    if a>100{
        return 0
    }else if a>25{
        return 1
    }else if a>1{
        return 5
    }else{
        return 10
    }
}
