package src

import "errors"

type Note struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	TimeCreated string `json:"createAt"`
	Favourite   bool   `json:"favourite"`
	User        *User
}

// admin
func (note *Note) All() (notes []Note, err error) {
	rows, err := Db.Query("Select id, title, body, created_at from notes")
	if err != nil {
		return
	}

	for rows.Next() {
		note := Note{}

		err = rows.Scan(&note.Id, &note.Title, &note.Body, &note.TimeCreated)
		if err != nil {
			return
		}

		notes = append(notes, note)
	}

	return
}

//
func (note *Note) Create() (err error) {
	if note.User == nil {
		return errors.New("user not found...")
	}

	statement := "insert into notes (title, body, user_id) values ($1, $2, $3)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(note.Title, note.Body, note.User.Id).Scan(&note.Id)
	if err != nil {
		return
	}

	return
}

//
func Notes(id int, userId int64) (note Note, err error) {
	err = Db.QueryRow("select id, title, body, created_at, favourite from notes where id=$1 and user_id=$2", id, userId).Scan(&note.Id, &note.Title, &note.Body, &note.TimeCreated, &note.Favourite)
	return
}

//
func (note *Note) Update() (err error) {
	_, err = Db.Exec("update notes set title = $2, body = $3 where id=$1 and user_id=$4", note.Id, note.Title, note.Body, note.User.Id)
	return
}

//
func (note *Note) Delete() (err error) {
	_, err = Db.Exec("delete from notes where id=$1 and user_id=$2", note.Id, note.User.Id)
	return
}

//
func Search(keyword string, userId int64) (notes []Note, err error) {
	rows, err := Db.Query("SELECT id, title, body, timeCreated, favourite FROM notes WHERE body LIKE $2 OR title LIKE $2 AND user_id = $1';", userId, keyword)

	if err != nil {
		return
	}

	for rows.Next() {
		note := Note{}

		err = rows.Scan(&note.Id, &note.Title, &note.Body, &note.TimeCreated, &note.Favourite)
		if err != nil {
			return
		}

		notes = append(notes, note)
	}

	return
}

//
func (note *Note) UpdateFavourite() (err error) {
	_, err = Db.Exec("update notes set favourite = $2 where id=$1 and user_id=$3", note.Id, note.Favourite, note.User.Id)
	return
}
