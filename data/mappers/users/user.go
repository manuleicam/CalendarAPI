package userMapper

import (
	userBD "calendarApi/data/models/users"
	. "calendarApi/module/users"
)

func UserToBdModel(user User) userBD.User {

	return userBD.User{
		Id:       user.GetId(),
		Name:     user.GetName(),
		Position: user.GetPosition(),
	}
}

func UserToModel(user userBD.User) User {

	return *NewUser(user.Id, user.Name, user.Position)
}
