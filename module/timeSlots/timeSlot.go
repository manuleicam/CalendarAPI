package timeSlots

import "time"

// DISP = AVAILABLE

// This is the ID Generator
var timeSlotSeq int = 0

type TimeSlot struct {
	userID int
	days   []DaysDisp
}
type DaysDisp struct {
	id      int
	day     time.Time
	hourBeg int
	hourEnd int
}

type TimeSlotJSON struct {
	UserID int            `json:"userID"`
	Days   []DaysDispJSON `json:"days"`
}
type DaysDispJSON struct {
	Id      int       `json:"id"`
	Day     time.Time `json:"day"`
	HourBeg int       `json:"hourBeg"`
	HourEnd int       `json:"hourEnd"`
}

func NewTimeSlot(userId int, daysDisps []DaysDisp) *TimeSlot {

	newDaysDisps := []DaysDisp{}

	for _, daysDisp := range daysDisps {
		newDaysDisps = append(newDaysDisps, *NewDaysDisp(daysDisp.id, daysDisp.day, daysDisp.hourBeg, daysDisp.hourEnd))
	}

	timeSlot := TimeSlot{userId, newDaysDisps}

	return &timeSlot
}

func NewDaysDisp(id int, day time.Time, hourBeg int, hourEnd int) *DaysDisp {
	return &DaysDisp{id, day, hourBeg, hourEnd}
}

func InitTimeSlot(userId int, day time.Time, hourBeg int, hourEnd int) TimeSlot {

	initDay := DaysDisp{0, day, hourBeg, hourEnd}
	DaysDisp := []DaysDisp{initDay}
	timeSlot := TimeSlot{userId, DaysDisp}

	return timeSlot

}

func (daysDis *DaysDisp) GetId() int {
	return daysDis.id
}

func (daysDis *DaysDisp) GetDay() time.Time {
	return daysDis.day
}

func (daysDis *DaysDisp) GetHourEnd() int {
	return daysDis.hourEnd
}

func (daysDis *DaysDisp) GetHourBeg() int {
	return daysDis.hourBeg
}

func (timeSlot *TimeSlot) GetUserID() int {
	return timeSlot.userID
}

func (timeSlot *TimeSlot) GetDays() []DaysDisp {

	return timeSlot.days

}

func (timeSlot *TimeSlot) SetUserID(id int) {

	timeSlot.userID = id

}

func (timeSlot *TimeSlot) setNewDay(newDay time.Time, hourBeg int, hourEnd int) int {

	for _, day := range timeSlot.days {

		if day.day == newDay {
			//hoursDisp := HoursDisp{hourBeg, hourEnd}

			if day.hourBeg == hourBeg {
				//timeSlot.days[dayIndex].hours = append(timeSlot.days[dayIndex].hours, hoursDisp)
				//day.hours = append(day.hours)
				return -1
			}
		}

	}

	newDayDisp := DaysDisp{timeSlotSeq, newDay, hourBeg, hourEnd}
	timeSlotSeq++

	timeSlot.days = append(timeSlot.days, newDayDisp)

	return 1

}

func (timeSlot *TimeSlot) CheckIfHoursExist(daysDisp DaysDisp) bool {

	for _, day := range timeSlot.days {

		if day.day == daysDisp.day && day.hourBeg == daysDisp.hourBeg && day.hourEnd == daysDisp.hourEnd {
			return true
		}

	}
	return false

}

// this will take a normal struct and convert it to a JSON struct
// The only difference is that all the camps in the JSON struct are public so that can be used in the json package
func TimeSlotToJSON(timeSlot TimeSlot) TimeSlotJSON {

	//stringDay := timeSlot.day.Format("2000-01-01 00:00:00")
	//stringDay := timeSlot.day.String()
	daysDispList := make([]DaysDispJSON, 0)
	dayDisp := timeSlot.days
	for _, day := range dayDisp {
		daysDispList = append(daysDispList, DaysDispJSON{day.id, day.day, day.hourBeg, day.hourEnd})
	}

	return TimeSlotJSON{timeSlot.userID, daysDispList}

}

// this will take a JSON struct and convert it to a normal struct
// The only difference is that all the camps in the JSON struct are public so that can be used in the json package
func JSONToTimeSlot(timeSlot TimeSlotJSON) TimeSlot {

	//stringDay := timeSlot.day.Format("2000-01-01 00:00:00")
	//stringDay := timeSlot.day.String()
	daysDispList := make([]DaysDisp, 0)
	dayDisp := timeSlot.Days
	for _, day := range dayDisp {
		daysDispList = append(daysDispList, DaysDisp{day.Id, day.Day, day.HourBeg, day.HourEnd})
	}

	return TimeSlot{timeSlot.UserID, daysDispList}

}

// this will take a list of normal struct and convert it to a list of JSON struct
// The only difference is that all the camps in the JSON struct are public so that can be used in the json package
func TimeSlotListToJSON(timeSlotList []TimeSlot) []TimeSlotJSON {

	timeSlotListJSON := make([]TimeSlotJSON, 0)

	for _, timeSlot := range timeSlotList {

		timeSlotJSON := TimeSlotToJSON(timeSlot)

		timeSlotListJSON = append(timeSlotListJSON, timeSlotJSON)

	}
	return timeSlotListJSON

}

// this will take a normal struct and convert it to a JSON struct
// The only difference is that all the camps in the JSON struct are public so that can be used in the json package
func DayDistToJSON(dayDisp []DaysDisp) []DaysDispJSON {

	daysDispListJSON := make([]DaysDispJSON, 0)

	for _, day := range dayDisp {

		daysDispJSON := DaysDispJSON{0, day.day, day.hourBeg, day.hourEnd}
		daysDispListJSON = append(daysDispListJSON, daysDispJSON)

	}
	return daysDispListJSON
}
