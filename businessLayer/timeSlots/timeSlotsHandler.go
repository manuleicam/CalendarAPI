package timeSlotsHandler

import (
	. "calendarApi/module/timeSlots"
	. "calendarApi/repositories/timeSlots"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TimeSlotsUserMatchHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getUsersTimeSlotMatch(w, r)
	default:
		w.WriteHeader(404)
	}

}

func TimeSlotUserHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getTimeSlotUser(w, r)
	case "POST":
		createTimeSlotUser(w, r)
	default:
		w.WriteHeader(404)
	}

}

func TimeSlotDeleteHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "DELETE":
		deleteUserTimeSlot(w, r)
	default:
		w.WriteHeader(404)
	}

}

func getUsersTimeSlotMatch(w http.ResponseWriter, r *http.Request) {

	userIds := r.URL.Query().Get("id")
	usersList := make([]int, 0)

	// Transforme the URL query into an array, and ignore everything that is not a number
	for _, id := range userIds {

		newId := string(id)

		if newId >= "0" && newId <= "9" {
			newIdInt, _ := strconv.Atoi(newId)
			usersList = append(usersList, newIdInt)
		}

	}

	// Go user by user and check which time slots match between them
	// Saves those timeslots in a slice and use the slice to compare the time slots with other users
	// In the end only the timeslots matching between all users will still be in the slice
	daysDisp := make([]DaysDisp, 0)
	for _, userId := range usersList {

		daysDisp = GetMatchTimeSlots(userId, daysDisp)

		// If the array is empty after the last function means that there is no match between users, so no need to look for more users
		if len(daysDisp) == 0 {
			break
		}

	}

	daysDispListJSON := DayDistToJSON(daysDisp)
	daysDispJSON, _ := json.Marshal(daysDispListJSON)

	w.WriteHeader(http.StatusOK)

	w.Write(daysDispJSON)

}

func deleteUserTimeSlot(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	userId, _ := vars["userId"]
	timeSlotid, _ := vars["timeSlotId"]

	userIdInt, userErr := strconv.Atoi(userId)
	timeSlotidInt, tsErr := strconv.Atoi(timeSlotid)

	if userErr != nil || tsErr != nil {
		w.WriteHeader(422)

		msg, _ := json.Marshal("Wrong ID type")

		w.Write(msg)

		return
	}

	res := DeleteUserTimeSlotBD(userIdInt, timeSlotidInt)

	switch res {
	case -1:
		w.WriteHeader(http.StatusNotFound)

		msg, _ := json.Marshal("User not found")

		w.Write(msg)
	case 0:
		w.WriteHeader(http.StatusNotFound)

		msg, _ := json.Marshal("Time slot not found")

		w.Write(msg)
	case 1:
		w.WriteHeader(http.StatusOK)

		msg, _ := json.Marshal("Time slot deleted")

		w.Write(msg)
	}

}

func createTimeSlotUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := vars["id"]

	idInt, idConvertError := strconv.Atoi(id)

	if idConvertError != nil {
		w.WriteHeader(422)

		msg, _ := json.Marshal("Wrong ID type")

		w.Write(msg)

		return
	}

	body, err := ioutil.ReadAll(r.Body)
	var timeSlotJSON TimeSlotJSON
	err = json.Unmarshal(body, &timeSlotJSON)

	if err != nil {

		w.WriteHeader(422)

		msg, _ := json.Marshal("One or more fields are wrong")

		w.Write(msg)

		return

	}

	timeSlotJSON.UserID = idInt

	newTimeSlot := JSONToTimeSlot(timeSlotJSON)

	res := AddNewTimeSlotBD(newTimeSlot)

	if res == 1 {

		msg, _ := json.Marshal("Time Slot Created")

		w.WriteHeader(http.StatusCreated)
		w.Write(msg)

	}

}

func getTimeSlotUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := vars["id"]

	idInt, idConvertError := strconv.Atoi(id)

	if idConvertError != nil {
		w.WriteHeader(422)

		msg, _ := json.Marshal("Wrong ID type")

		w.Write(msg)

		return
	}

	timeSlots, foundUser := GetUserTimeSlotBD(idInt)

	if !foundUser {

		msg, _ := json.Marshal("User not found")

		w.WriteHeader(404)
		w.Write(msg)

	} else {

		timeSlotsStructJSON := TimeSlotToJSON(timeSlots)
		timeSlotsJSON, _ := json.Marshal(timeSlotsStructJSON)

		w.WriteHeader(http.StatusOK)

		w.Write(timeSlotsJSON)

	}

}
