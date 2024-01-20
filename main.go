package main

import (
	. "calendarApi/repository/timeSlots"
	. "calendarApi/repository/users"
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
