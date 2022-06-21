package usersDB

import (
	. "calendarApi/module/users"
	"fmt"
)

var userTable map[int]User

func StartUsersDB() {

	userTable = make(map[int]User)
	//userTable := make(map[int]User)

	userTable[1] = IniUser(1, "Inês", "Interviewer")
	userTable[2] = IniUser(2, "Ingrid", "Interviewer")
	userTable[3] = IniUser(3, "Carl", "Candidate")

}

func DeleteAllUsersBD() {
	userTable = make(map[int]User)
}

func GetAllUsersBD() []User {

	allUsers := make([]User, 0)

	for _, user := range userTable {

		allUsers = append(allUsers, user)

	}

	return (allUsers)

}

func GetOneUserBD(userID int) (User, bool) {

	userFound, error := userTable[userID]

	return userFound, error

}

func DeleteOneUserBD(userID int) bool {

	_, userFound := userTable[userID]

	if userFound {

		delete(userTable, userID)

		return true
	}

	return false

}

func UpdateOneUserBD(id int, newUser UserJSON) int {

	oldUser, userFound := userTable[id]

	if userFound {

		oldUser.SetName(newUser.Name)
		oldUser.SetPosition(newUser.Position)

		userTable[id] = oldUser

		return id
	}

	return -1

}

func getMaxID() int {

	return len(userTable) + 1

}

func AddUser(name string, pos string) int {

	id := getMaxID()

	user := IniUser(id, name, pos)

	userTable[id] = user

	return id

}

func PrintUserBD() {

	fmt.Println("USERS: ", userTable)

}
