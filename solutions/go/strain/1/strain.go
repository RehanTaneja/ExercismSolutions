package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](c []T, f func(T)bool)[]T{
    var a []T
    for _,x := range c{
        if f(x){
            a = append(a,x)
        }
    }
    return a
}

func Discard[T any](c []T, f func(T)bool)[]T{
    var a []T
    for _,x := range c{
        if !f(x){
            a = append(a,x)
        }
    }
    return a
}
