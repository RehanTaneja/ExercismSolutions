package complexnumbers

import "math"

// Define the Number type here.

type Number struct{
    a float64
    b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	var r float64 = n1.a + n2.a
    var i float64 = n1.b + n2.b
    return Number{
        a : r,
        b : i,
    }
}

func (n1 Number) Subtract(n2 Number) Number {
	var r float64 = n1.a - n2.a
    var i float64 = n1.b - n2.b
    return Number{
        a : r,
        b : i,
    }
}

func (n1 Number) Multiply(n2 Number) Number {
	var r float64 = (n1.a*n2.a) - (n1.b*n2.b)
    var i float64 = (n1.b*n2.a) + (n1.a*n2.b)
    return Number{
        a : r,
        b : i,
    }
}

func (n Number) Times(factor float64) Number {
    return Number{
        a : n.a*factor,
        b : n.b*factor,
    }
}

func (n1 Number) Divide(n2 Number) Number {
	var r float64 = ((n1.a*n2.a) + (n1.b*n2.b)) / ((n2.a*n2.a) + (n2.b*n2.b))
    var i float64 = ((n1.b*n2.a) - (n1.a*n2.b)) / ((n2.a*n2.a) + (n2.b*n2.b))
    return Number{
        a : r,
        b : i,
    }
}

func (n Number) Conjugate() Number {
    return Number{
        a : n.a,
        b : -n.b,
    }
}

func (n Number) Abs() float64 {
	return math.Sqrt((n.a*n.a) + (n.b*n.b))
}

func (n Number) Exp() Number {
	var r float64 = math.Exp(n.a) * math.Cos(n.b)
    var i float64 = math.Exp(n.a) * math.Sin(n.b)
    return Number{
        a : r,
        b : i,
    }
}
