package n4u_admin

// the structure of the notes as it is in the api
// makes it easy for parsing json
type Note struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	TimeCreated string `json:"createAt"`
}

type NotesTable struct {
	Notes               []Note
	Name                string
	LastRequestResponse string
}
