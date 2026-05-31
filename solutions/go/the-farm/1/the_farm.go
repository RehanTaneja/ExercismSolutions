package thefarm

import "errors"
import "fmt"

// TODO: define the 'DivideFood' function 
func DivideFood(fc FodderCalculator, numCows int) (float64,error){
    total_fodder_amount,err := fc.FodderAmount(numCows)
    factor,err_two := fc.FatteningFactor()
    if err == nil && err_two == nil{
        return (total_fodder_amount*factor)/float64(numCows),err
    }else{
        if err_two == nil{
            return 0.0,err
        }else if err == nil{
            return 0.0, err_two
        }
    }
    return 0.0, err
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64,error) {
    if numCows>0{
        return DivideFood(fc, numCows)
    }else{
        var ErrNotFound = errors.New("invalid number of cows")
        return 0.0, ErrNotFound
    }
}

type InvalidCowsError struct{
    message string
}

func (i *InvalidCowsError) Error() string{
    return i.message
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error{
    if numCows<0{
        var msg = fmt.Sprintf("%v cows are invalid: there are no negative cows",numCows)
        return &InvalidCowsError{message: msg}
    }else if numCows == 0{
        return &InvalidCowsError{message: "0 cows are invalid: no cows don't need food"}
    }else{
        return nil
    }
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
