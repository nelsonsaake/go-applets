package main

import (
	"projects/saelections/pkg/sysout"
)

func RepoCategoryByHref(href string) (c Category, err error) {
	err = sqlite.Where(&Category{Href: href}).First(&c).Error
	return
}

func RepoCategoryID(href string) (int64, error) {
	c, err := RepoCategoryByHref(href)
	return c.ID, err
}

func RepoCategorySave(c *Category) (err error) {
	categories, products, v := c.Categories, c.Products, sysout.Verbose
	c.Categories, c.Products = nil, nil

	if x, err := RepoCategoryByHref(c.Href); err == nil {
		c.ID = x.ID
		if c.CategoryID == 0 {
			c.CategoryID = x.CategoryID
		}
	}

	if err := sqlite.Save(c).Error; err != nil {
		return err
	}

	for _, subcategory := range categories {
		subcategory.CategoryID = c.ID
	}

	if err := RepoCategorySaveAll(categories); err != nil {
		v(err)
	}

	if err := RepoProductSaveAll(products); err != nil {
		v(err)
	}

	err = sqlite.Model(c).Association("Products").Append(products)

	c.Categories, c.Products = categories, products

	return
}

func RepoCategorySaveAll(categories []*Category) (err error) {
	for _, c := range categories {
		if err := RepoCategorySave(c); err != nil {
			return err
		}
	}
	return
}

func RepoCategoriesCreateFromMap(categories map[string]*Category) (err error) {
	for _, cat := range categories {
		if err = RepoCategorySave(cat); err != nil {
			return err
		}
	}
	return
}
