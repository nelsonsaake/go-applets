package main

type Product struct {
	Model
	Name            string   `json:"name" gorm:"UniqueIndex;not null"`
	Price           string   `json:"price"`
	Description     string   `json:"description"`
	FullDescription string   `json:"fullDescription"`
	CategoryHref    string   `json:"categoryHref"`
	CategoryName    string   `json:"categoryName"`
	Href            string   `json:"href" gorm:"UniqueIndex;not null"`
	Tags            []*Tag   `json:"tags,omitempty" gorm:"many2many:product_tags;" validate:"-" faker:"-"`
	Images          []*Image `json:"images,omitempty" validate:"-" faker:"-"`
}

func init() {
	modelstomigrate = append(modelstomigrate, &Product{})
}
