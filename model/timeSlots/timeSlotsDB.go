package timeSlotsDB

import (
	. "calendarApi/module/timeSlots"
	"fmt"
	"time"
)

var timeSlotsTable map[int]TimeSlot

// Functions that populates my DB based on the example cases used in GITLAB
// Sunday = 0
func StartTimeSlotsDB() {

	timeSlotsTable = make(map[int]TimeSlot)

	todayDate := time.Now()
	year := todayDate.Year()
	month := todayDate.Month()
	day := todayDate.Day()
	timeEX := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	newTimeSlot := TimeSlot{} //InitTimeSlot(1, timeEX, 9, 16)
	newTimeSlot.SetUserID(1)

	for i := 0; i < 7; i++ {

		insertDate := timeEX.AddDate(0, 0, i)
		wd := insertDate.Weekday()

		if wd == 0 || wd == 6 {
			continue
		}

		newTimeSlot.SetNewDay(insertDate, 9, 16)

	}

	timeSlotsTable[1] = newTimeSlot

	newTimeSlot = TimeSlot{} // newTimeSlot = InitTimeSlot(2, timeEX, 12, 18)
	newTimeSlot.SetUserID(2)

	for i := 0; i < 7; i++ {
		insertDate := timeEX.AddDate(0, 0, i)
		wd := insertDate.Weekday()

		if wd == 0 || wd == 6 || wd == 5 {
			continue
		}

		if wd == 2 || wd == 4 {

			newTimeSlot.SetNewDay(insertDate, 9, 12)

		} else {

			newTimeSlot.SetNewDay(insertDate, 12, 18)

		}

	}

	timeSlotsTable[2] = newTimeSlot

	newTimeSlot = TimeSlot{}
	newTimeSlot.SetUserID(3)

	for i := 0; i < 7; i++ {
		insertDate := timeEX.AddDate(0, 0, i)
		wd := insertDate.Weekday()

		if wd == 0 || wd == 6 {
			continue
		}

		if wd == 3 {
			newTimeSlot.SetNewDay(insertDate, 10, 12)
			//newTimeSlot := InitTimeSlot(3, insertDate, 9, 10)
			//timeSlotsTable[3] = append(timeSlotsTable[3], newTimeSlot)
		}

		newTimeSlot.SetNewDay(insertDate, 9, 10)
		//newTimeSlot := InitTimeSlot(3, insertDate, 10, 12)

		//timeSlotsTable[3] = append(timeSlotsTable[3], newTimeSlot)

	}

	timeSlotsTable[3] = newTimeSlot

}

// return -1 if no user was found
// return 0 if the Time slot was not found on the user
// return 1 if the time slot was deleted
func DeleteUserTimeSlotBD(userId, timeSlotid int) int {

	userTimeSlot, foundUser := timeSlotsTable[userId]

	if foundUser {

		res := userTimeSlot.DeleteTimeSlot(timeSlotid)
		timeSlotsTable[userId] = userTimeSlot

		return res

	}

	return -1

}

func AddNewTimeSlotBD(timeSlot TimeSlot) int {

	userTimeSlot, foundUser := timeSlotsTable[timeSlot.GetUserID()]

	if foundUser {

		dayList := timeSlot.GetDays()

		for _, day := range dayList {

			userTimeSlot.SetNewDay(day.GetDay(), day.GetHourBeg(), day.GetHourEnd())

		}

		timeSlotsTable[timeSlot.GetUserID()] = userTimeSlot

		return 1

	} else {

		timeSlotsTable[timeSlot.GetUserID()] = timeSlot

		return 1

	}

}

func GetMatchTimeSlots(userId int, timeSlots []DaysDisp) []DaysDisp {

	daysDisp := make([]DaysDisp, 0)
	userTimeSlot, foundUser := timeSlotsTable[userId]

	if foundUser {

		if len(timeSlots) == 0 {

			return userTimeSlot.GetDays()

		} else {

			for _, matchTime := range timeSlots {

				if userTimeSlot.CheckIfHoursExist(matchTime) {

					daysDisp = append(daysDisp, matchTime)

				}

			}

		}

	}

	return daysDisp

}

// returns an empty struct and FALSE if the user was not found
func GetUserTimeSlotBD(userID int) (TimeSlot, bool) {

	timeSlot, foundUser := timeSlotsTable[userID]

	return timeSlot, foundUser

}

func PrintTimeSlotsBD() {

	fmt.Println("TimeSlots: ", timeSlotsTable)

}
