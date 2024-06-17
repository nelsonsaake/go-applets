package main

type Image struct {
	Model
	Url       string `json:"url" gorm:"UniqueIndex;not null"`
	Local     string `json:"local"`
	ProductID int64  `json:"productId,omitempty"`
}

func init() {
	modelstomigrate = append(modelstomigrate, &Image{})
}
