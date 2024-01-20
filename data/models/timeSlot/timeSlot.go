package timeSlot

type TimeSlot struct {
	UserID   int
	Meetings []Meeting
}

func (timeSlot *TimeSlot) DeleteTimeSlot(timeSlotId int) {

	for tsId, day := range timeSlot.Meetings {

		if day.Id == timeSlotId {

			timeSlot.Meetings = append(timeSlot.Meetings[:tsId], timeSlot.Meetings[tsId+1:]...)

			return
		}
	}
}

func (timeSlot *TimeSlot) AddMeetingInfo(meeting Meeting) {

	timeSlot.Meetings = append(timeSlot.Meetings, meeting)
}
