package timeSlotMapper

import (
	timeSlotBD "calendarApi/data/models/timeSlot"
	. "calendarApi/module/timeSlots"
)

func MeetingToBdModel(daysDisp DaysDisp) timeSlotBD.Meeting {

	return timeSlotBD.Meeting{
		Id:      daysDisp.GetId(),
		Day:     daysDisp.GetDay(),
		HourBeg: daysDisp.GetHourBeg(),
		HourEnd: daysDisp.GetHourEnd(),
	}
}

func MeetingsToBdModel(daysDisps []DaysDisp) []timeSlotBD.Meeting {
	var meetings = make([]timeSlotBD.Meeting, 0)

	for _, daysDisp := range daysDisps {
		meetings = append(meetings, MeetingToBdModel(daysDisp))
	}

	return meetings
}

func MeetingToModel(meeting timeSlotBD.Meeting) DaysDisp {

	return *NewDaysDisp(meeting.Id, meeting.Day, meeting.HourBeg, meeting.HourEnd)
}

func MeetingsToModel(meetingsBD []timeSlotBD.Meeting) []DaysDisp {
	var daysDisps = make([]DaysDisp, 0)

	for _, meetingBD := range meetingsBD {
		daysDisps = append(daysDisps, MeetingToModel(meetingBD))
	}

	return daysDisps
}
