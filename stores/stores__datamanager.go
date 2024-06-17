package main

type DataManager struct {
	MainCategories map[string]*Category
	SubCategories  map[string]*Category
	Products       map[string]*Product
	Tags           map[string]*Tag
	Images         map[string]*Image
}

func NewDataManager() (dm *DataManager) {
	return &DataManager{
		MainCategories: map[string]*Category{},
		SubCategories:  map[string]*Category{},
		Products:       map[string]*Product{},
		Tags:           map[string]*Tag{},
		Images:         map[string]*Image{},
	}
}

func NewDataManageFromJsonFile(file string) (dm *DataManager, err error) {
	dm = NewDataManager()
	if err = loadJsonFile(file, &dm.MainCategories); err != nil {
		return
	}
	for _, c := range dm.MainCategories {
		for _, sc := range c.Categories {
			dm.SubCategories[sc.Href] = sc
		}
	}
	for _, sc := range dm.SubCategories {
		for _, p := range sc.Products {
			dm.Products[p.Href] = p
		}
	}
	for _, p := range dm.Products {
		for _, t := range p.Tags {
			dm.Tags[t.Text] = t
		}
	}
	for _, p := range dm.Products {
		for _, img := range p.Images {
			dm.Images[img.Url] = img
		}
	}
	return
}
