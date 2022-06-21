package main

import (
	. "calendarApi/model/timeSlots"
	. "calendarApi/model/users"
	. "calendarApi/routes"
)

func main() {

	// This 2 function start my DB
	StartUsersDB()
	StartTimeSlotsDB()

	HandleRequests()

	//PrintUserBD()
	//PrintTimeSlotsBD()

}
