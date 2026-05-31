package sorting

import "fmt"
import "strconv"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f",f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	var myfloat float64 = float64(nb.Number())
    return fmt.Sprintf("This is a box containing the number %.1f",myfloat)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type){
        case FancyNumber:
        	var i, err = strconv.Atoi(fnb.Value())
        	if err == nil{
                return i
            }
        	return 0
        default:
        	return 0
    }
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fnb.(type){
        case FancyNumber:
        	var i,err = strconv.Atoi(fnb.Value())
        	if err == nil{
                return fmt.Sprintf("This is a fancy box containing the number %.1f",float64(i))
            }else{
                return fmt.Sprintf("This is a fancy box containing the number 0.0")
            }
        default:
        	return fmt.Sprintf("This is a fancy box containing the number 0.0")
    }
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type){
        case int:
        	var myfloat float64 = float64(i.(int))
        	return DescribeNumber(myfloat)
        case float64:
        	return DescribeNumber(i.(float64))
        case NumberBox:
        	return DescribeNumberBox(i.(NumberBox))
        case FancyNumberBox:
        	return DescribeFancyNumberBox(i.(FancyNumberBox))
        default:
        	return "Return to sender"
    }
}
