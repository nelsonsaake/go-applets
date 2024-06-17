package main

type Category struct {
	Model
	Name       string      `json:"name" gorm:"UniqueIndex;not null"`
	Href       string      `json:"href" gorm:"UniqueIndex;not null"`
	ParentHref string      `json:"parentHref"`
	CategoryID int64       `json:"categoryId,omitempty"`
	Categories []*Category `json:"categories,omitempty" validate:"-" faker:"-"`
	Products   []*Product  `json:"products,omitempty" gorm:"many2many:category_products;" validate:"-" faker:"-"`
}

func init() {
	modelstomigrate = append(modelstomigrate, &Category{})
}
