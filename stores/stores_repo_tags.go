package main

func RepoTagByText(text string) (tag Tag, err error) {
	err = sqlite.Where(&Tag{Text: text}).First(&tag).Error
	return
}

func RepoTagID(text string) (int64, error) {
	tag, err := RepoTagByText(text)
	return tag.ID, err
}

func RepoTagSave(tag *Tag) (err error) {
	if id, err := RepoTagID(tag.Text); err == nil {
		tag.ID = id
	}
	return sqlite.Save(tag).Error
}

func RepoTagSaveAll(tags []*Tag) (err error) {
	for _, tag := range tags {
		if err := RepoTagSave(tag); err != nil {
			return err
		}
	}
	return nil
}
