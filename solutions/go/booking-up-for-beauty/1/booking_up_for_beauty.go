package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    //layout := "January 2, 2006 15:04:05"
	layout :="1/2/2006 15:04:05"
    t,err := time.Parse(layout,date)
    if err!=nil{
        fmt.Println(err)
    }
    fmt.Println(t)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
	t,err := time.Parse(layout,date)
    if err!=nil{
        fmt.Println(err)
    }
    var hasPassed bool = t.Before(time.Now())
    fmt.Println(t,time.Now())
    return hasPassed
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
	t,err := time.Parse(layout,date)
    if err!=nil{
        fmt.Println(err)
    }
    if t.Hour()>=12 && t.Hour()<18{
        return true
    }else{
        return false
    }
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
	t,err := time.Parse(layout,date)
    if err!=nil{
        fmt.Println(err)
    }
    var s = fmt.Sprintf("You have an appointment on " +t.Format("Monday, January 2, 2006, at 15:04")+".")
    return s
}	

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    date := "2025-09-15 00:00:00 +0000 UTC"
    layout := "2006-01-02 15:04:05 -0700 MST"
	t,err := time.Parse(layout,date)
    if err!=nil{
        fmt.Println(err)
    }
    return t
}
