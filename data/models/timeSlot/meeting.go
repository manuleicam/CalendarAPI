package timeSlot

import "time"

type Meeting struct {
	Id      int
	Day     time.Time
	HourBeg int
	HourEnd int
}
