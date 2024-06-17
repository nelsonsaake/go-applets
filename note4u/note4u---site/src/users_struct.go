package n4u_admin

// the structure of the Users as it is in the api
// makes it easy for parsing json

// the struct represents the users-structure as it is in at the api
type User struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	TimeCreated string `json:"createAt"`
	Notes       []Note
}

// the usersTable structure hold all information
// that will be required by the template to build the page
type UsersTable struct {
	Users               []User
	Name                string
	LastRequestResponse string
}
