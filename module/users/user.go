package users

type User struct {
	id       int
	name     string
	position string
}

type UserJSON struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

func NewUser(id int, name string, pos string) *User {
	return &User{id, name, pos}
}

func IniUser(id int, name string, pos string) User {

	newUser := User{id, name, pos}

	return newUser

}

func (user *User) GetId() int {
	return user.id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) GetPosition() string {
	return user.position
}

func (user *User) SetPosition(pos string) {
	user.position = pos
}

func (user *User) SetName(n string) {
	user.name = n
}

// this will take a normal struct and convert it to a JSON struct
// The only difference is that all the camps in the JSON struct are public so that can be used in the json package
func UserToJSON(user User) UserJSON {

	userJSON := UserJSON{user.id, user.name, user.position}

	return userJSON

}

// this will take a normal struct and convert it to a JSON struct
// The only difference is that all the camps in the JSON struct are public so that can be used in the json package
func UserListToJSON(userList []User) []UserJSON {

	userListToJson := make([]UserJSON, 0)

	for _, user := range userList {

		userListToJson = append(userListToJson, UserToJSON(user))

	}

	return userListToJson

}
