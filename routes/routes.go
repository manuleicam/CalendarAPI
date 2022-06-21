package routes

import (
	. "calendarApi/businessLayer/timeSlots"
	. "calendarApi/businessLayer/users"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	r := mux.NewRouter()

	/// User routes
	// Gets users, delete all users, creates new user
	r.HandleFunc("/users", UserHandler)
	// get, update and delete user info by user ID
	r.HandleFunc("/users/{id}", UserIDHandler)
	/// Times Slots routes
	// delete timeslot based on userid and timeslotid
	r.HandleFunc("/users/{userId}/time-slots/{timeSlotId}", TimeSlotDeleteHandler)
	// get and create user timeslot by user id
	r.HandleFunc("/time-slots/users/{id}", TimeSlotUserHandler)
	// get the match between users IDs (user id in query parameters)
	r.HandleFunc("/time-slots/users", TimeSlotsUserMatchHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
