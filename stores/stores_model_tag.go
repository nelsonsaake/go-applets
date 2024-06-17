package main

type Tag struct {
	Model
	Text string `json:"text" gorm:"UniqueIndex;not null"`
}

func init() {
	modelstomigrate = append(modelstomigrate, &Tag{})
}
