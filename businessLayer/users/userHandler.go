package userHandler

import (
	. "calendarApi/module/users"
	. "calendarApi/repository/users"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getAllUsersMethod(w, r)
	case "POST":
		createUserMethod(w, r)
	case "DELETE":
		deleteAllUsers(w, r)
	default:
		w.WriteHeader(404)
	}

}

func UserIDHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getOneUser(w, r)
	case "PUT":
		updateOneUser(w, r)
	case "DELETE":
		deleteOneUser(w, r)
	default:
		w.WriteHeader(404)
	}

}

func updateOneUser(w http.ResponseWriter, r *http.Request) {

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
	var userJSON UserJSON
	err = json.Unmarshal(body, &userJSON)

	if err != nil || userJSON.Name == "" || userJSON.Position == "" {

		w.WriteHeader(422)

		msg, _ := json.Marshal("One or more fields are wrong")

		w.Write(msg)

		return

	}

	errUpdate := UpdateOneUserBD(idInt, userJSON)

	if errUpdate == -1 {

		msg, _ := json.Marshal("User not found")

		w.WriteHeader(404)
		w.Write(msg)

	} else {

		w.WriteHeader(http.StatusOK)
		msg, _ := json.Marshal("ID: " + id + " Updated")

		w.Write(msg)

	}

}

func deleteAllUsers(w http.ResponseWriter, r *http.Request) {

	DeleteAllUsersBD()

	w.WriteHeader(http.StatusOK)
	msg, _ := json.Marshal("Users Removed")

	w.Write(msg)

}

func deleteOneUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, _ := vars["id"]

	idInt, idConvertError := strconv.Atoi(id)

	if idConvertError != nil {
		w.WriteHeader(422)

		msg, _ := json.Marshal("Wrong ID type")

		w.Write(msg)

		return
	}

	delSucc := DeleteOneUserBD(idInt)

	if delSucc {

		w.WriteHeader(http.StatusOK)
		msg, _ := json.Marshal("ID: " + id + " Deleted")

		w.Write(msg)

	} else {

		msg, _ := json.Marshal("User not found")

		w.WriteHeader(404)
		w.Write(msg)

	}

}

func createUserMethod(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	var userJSON UserJSON
	err = json.Unmarshal(body, &userJSON)

	if err != nil || userJSON.Name == "" || userJSON.Position == "" {

		w.WriteHeader(422)

		msg, _ := json.Marshal("One or more fields are wrong")

		w.Write(msg)

		return

	}

	newId := AddUser(userJSON.Name, userJSON.Position)

	w.WriteHeader(http.StatusCreated)
	msg, _ := json.Marshal("New Id: " + strconv.Itoa(newId))

	w.Write(msg)

}

func getOneUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, _ := vars["id"]

	idInt, idConvertError := strconv.Atoi(id)

	if idConvertError != nil {
		w.WriteHeader(422)

		msg, _ := json.Marshal("Wrong ID type")

		w.Write(msg)

		return
	}

	userFound, exists := GetOneUserBD(idInt)

	if exists {

		userStructJSON := UserToJSON(userFound)

		userJSON, _ := json.Marshal(userStructJSON)

		w.Write(userJSON)

	} else {

		msg, _ := json.Marshal("User not found")

		w.WriteHeader(404)
		w.Write(msg)

	}

}

func getAllUsersMethod(w http.ResponseWriter, r *http.Request) {

	userList := getAllUsers()
	w.Header().Set("Content-Type", "application/json")

	//fmt.Println(userList)

	userListJSON := UserListToJSON(userList)

	userJSON, _ := json.Marshal(userListJSON)

	w.Write(userJSON)

}

func getAllUsers() []User {

	return GetAllUsersBD()

}
