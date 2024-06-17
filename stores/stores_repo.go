package main

import (
	"fmt"
	"projects/saelections/app/repo/sqlitegorm"
	"projects/saelections/pkg/sysout"

	"gorm.io/gorm"
)

var sqlite *gorm.DB

func init() {
	if err := initrepo(); err != nil {
		sysout.Fatal("init repo failed: %v", err)
	}
}

func initrepo() (err error) {
	if sqlite, err = sqlitegorm.New(repofile); err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}
	if err = autoMigration(); err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}
	return
}

func autoMigration() error {
	return sqlite.AutoMigrate(modelstomigrate...)
}
