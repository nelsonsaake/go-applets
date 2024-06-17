package main

import "projects/saelections/pkg/str"

func RepoCourseSave(url string, setters ...func(c *Course)) (c *Course, err error) {
	c = &Course{URL: url}

	if !str.Empty(url) {
		sqlite.First(c) // ignoring errors, don't matter
	}

	for _, setter := range setters {
		setter(c)
	}

	if err = sqlite.Save(c).Error; err != nil {
		return
	}

	return
}
