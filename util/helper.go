package util

import (
	"time"
)

// StringToInt convert's string to an integer
func StringToInt(s string) int64 {
	var i int64
	for _, c := range s {
		i = i*10 + int64(c-'0')
	}
	return i
}

// HumanDate returns a human-readable date
func HumanDate(t time.Time) string {
	return t.Format("2006-01-02")
}
