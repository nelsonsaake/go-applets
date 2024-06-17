package main

func RepoDraftFromFile(file string) (err error) {
	dm, err := NewDataManageFromJsonFile(file)
	if err != nil {
		return
	}
	if err = RepoCategoriesCreateFromMap(dm.MainCategories); err != nil {
		return
	}
	return
}
