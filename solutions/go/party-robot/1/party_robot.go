package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	var s = fmt.Sprintf("Welcome to my party, %v!",name)
    return s
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	var s = fmt.Sprintf("Happy birthday %s! You are now %v years old!",name,age)
    return s
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    var table_name string
    if len(fmt.Sprintf("%d",table))==1{
        table_name=fmt.Sprintf("00%d",table)
    }else if len(fmt.Sprintf("%d",table))==2{
        table_name=fmt.Sprintf("0%d",table)
    }else{
        table_name=fmt.Sprintf("%d",table)
    }
	var s = fmt.Sprintf("%v\nYou have been assigned to table %v. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v.",Welcome(name),table_name,direction,distance,neighbor)
    return s
}
