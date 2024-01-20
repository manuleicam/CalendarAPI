package timeSlotsDB

import (
	timeSlotBD "calendarApi/data/models/timeSlot"
	. "calendarApi/module/timeSlots"
	. "calendarApi/data/mappers/timeSlot"
	"errors"
	"fmt"
)

// The map key is the time slot id
var timeSlotsTable map[int]timeSlotBD.TimeSlot

func DeleteUserTimeSlot(userId, timeSlotId int) {
	userTimeSlot, _ := timeSlotsTable[userId]

	userTimeSlot.DeleteTimeSlot(timeSlotId)

	timeSlotsTable[userId] = userTimeSlot
}

func AddUserMeetingInfo(userId int, newMeeting timeSlotBD.Meeting) error {
	userTimeSlot, _ := timeSlotsTable[userId]

	for _, meeting := range userTimeSlot.Meetings {
		if newMeeting.Day == meeting.Day && newMeeting.HourBeg <= meeting.HourEnd && newMeeting.HourEnd >= meeting.HourBeg {
			return errors.New("Schedule Conflits!")
		}
	}

	userTimeSlot.AddMeetingInfo(newMeeting)

	timeSlotsTable[userId] = userTimeSlot

	return nil
}

func GetUserMeetingsBD(userID int) []DaysDisp {
	meetings, _ := timeSlotsTable[userID]

	return MeetingsToModel(meetings.Meetings)
}

func PrintTimeSlotsBD() {

	fmt.Println("TimeSlots: ", timeSlotsTable)
}