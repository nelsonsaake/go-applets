package main

import (
	"fmt"
	"projects/saelections/pkg/sysout"
	"strings"
	"unicode"
)

func clean(q string) string {
	for _, ch := range q {
		if !(unicode.IsDigit(ch) || unicode.IsLetter(ch)) {
			q = strings.ReplaceAll(q, string(ch), " ")
		}
	}
	for {
		q = strings.ReplaceAll(q, "  ", " ")
		if !strings.Contains(q, "  ") {
			break
		}
	}
	q = strings.ToLower(q)
	q = strings.ReplaceAll(q, " ", "%")
	return q
}

func like(col, q string) string {
	q = clean(q)
	q = fmt.Sprintf("lower(%s) like '%s'", col, "%"+q+"%")
	return q
}

func RepoProductsFind(q string) (products []Product, err error) {
	products = []Product{}
	err = sqlite.
		Model(&Product{}).
		Preload("Tags").
		Preload("Images").
		Where(like("name", q)).
		Or(like("description", q)).
		Find(&products).Error
	return
}

func RepoProductFind(q string) (product Product, err error) {
	err = sqlite.
		Model(&Product{}).
		Preload("Tags").
		Preload("Images").
		Where(like("name", q)).
		Or(like("description", q)).
		First(&product).Error
	return
}

func RepoProductsWithNoImages() (products []Product, err error) {
	sql := `
		SELECT *
		FROM products 
		WHERE id NOT in
		(
			SELECT DISTINCT product_id
			FROM images
		);
	`
	err = sqlite.
		Model(&Image{}).
		Raw(sql).
		Find(&products).
		Error
	return
}

func RepoProductByHref(href string) (p Product, err error) {
	err = sqlite.Where(&Product{Href: href}).First(&p).Error
	return
}

func RepoProductID(href string) (int64, error) {
	p, err := RepoProductByHref(href)
	return p.ID, err
}

func RepoProductSave(p *Product) (err error) {
	images, tags, v := p.Images, p.Tags, sysout.Verbose
	p.Images, p.Tags = nil, nil

	if id, err := RepoProductID(p.Href); err == nil {
		p.ID = id
	}

	if err := sqlite.Save(p).Error; err != nil {
		return err
	}

	for _, image := range images {
		image.ProductID = p.ID
	}

	if err := RepoImageSaveAll(images); err != nil {
		v(err)
	}

	if err := RepoTagSaveAll(tags); err != nil {
		v(err)
	}

	if err := sqlite.Model(p).Association("Tags").Append(tags); err != nil {
		v(err)
	}

	p.Images, p.Tags = images, tags

	return
}

func RepoProductSaveAll(products []*Product) (err error) {
	for _, product := range products {
		if err := RepoProductSave(product); err != nil {
			return err
		}
	}
	return nil
}
