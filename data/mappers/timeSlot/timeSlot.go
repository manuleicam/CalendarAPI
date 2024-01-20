package timeSlotMapper

import (
	timeSlotBD "calendarApi/data/models/timeSlot"
	. "calendarApi/module/timeSlots"
)

func TimeSlotToBdModel(timeSlot TimeSlot) timeSlotBD.TimeSlot {

	return timeSlotBD.TimeSlot{
		UserID:   timeSlot.GetUserID(),
		Meetings: MeetingsToBdModel(timeSlot.GetDays()),
	}
}

func TimeSlotToModel(timeSlot timeSlotBD.TimeSlot) TimeSlot {

	return *NewTimeSlot(timeSlot.UserID, MeetingsToModel(timeSlot.Meetings))
}
