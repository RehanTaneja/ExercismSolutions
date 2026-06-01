package etl

import(
    "strings"
)

func Transform(in map[int][]string) map[string]int {
	a := make(map[string]int)
    for i,x := range in{
        for _,y := range x{
            a[strings.ToLower(string(y))] = i
        }
    }
    return a
}
