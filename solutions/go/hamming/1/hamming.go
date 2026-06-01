package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
	if len(a)!=len(b){
        return 0,errors.New("Length not matching")
    }
    var cnt int = 0
    for i,_ := range a {
        if a[i]!=b[i]{
            cnt++
        }
    }
    return cnt,nil
}
