package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    t1, err := time.Parse(layout, date)
    if err != nil {
        fmt.Println("Error parsing date:", err)
    }
	t2 := time.Now()
    if t2.After(t1) {
        return true
    }
    return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    t1, _ := time.Parse(layout, date)
    hour := t1.Hour()
    if hour >= 12 && hour <= 18 {
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t1, _ := time.Parse(layout, date)
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t1.Weekday(), t1.Month(), t1.Day(), t1.Year(), t1.Hour(), t1.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02 15:04:05 -0700 MST"
	str := "2025-09-15 00:00:00 +0000 UTC"
    t, _ := time.Parse(layout, str)
    return t
}
