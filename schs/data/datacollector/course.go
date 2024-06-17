package main

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	automigration = append(automigration, &Course{})
}

type Course struct {
	URL           string `gorm:"primarykey"`
	DataSourceURL string
	Name          string
	School        string
	ClassStyle    string
	Duration      string
	ProgramStyle  string
	LongReq       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func Name(name string) func(*Course) {
	return func(c *Course) {
		c.Name = name
	}
}

func URL(url string) func(*Course) {
	return func(c *Course) {
		c.URL = url
	}
}

func DataSourceURL(dsu string) func(*Course) {
	return func(c *Course) {
		c.DataSourceURL = dsu
	}
}

func School(school string) func(*Course) {
	return func(c *Course) {
		c.School = school
	}
}

func Duration(duration string) func(*Course) {
	return func(c *Course) {
		c.Duration = duration
	}
}

func ClassStyle(cs string) func(*Course) {
	return func(c *Course) {
		c.ClassStyle = cs
	}
}

func ProgramStyle(ps string) func(*Course) {
	return func(c *Course) {
		c.ProgramStyle = ps
	}
}

func LongReq(lq string) func(*Course) {
	return func(c *Course) {
		c.LongReq = lq
	}
}
