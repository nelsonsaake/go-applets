package main

import (
	"fmt"
	"projects/saelections/app/repo/sqlitegorm"
	"projects/saelections/pkg/sysout"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	sqlite        *gorm.DB
	repofile      = "_repo.db"
	automigration = []interface{}{}
)

func init() {
	if err := InitRepo(); err != nil {
		sysout.Fatal("init repo failed: %v", err)
	}
	SilentLogger()
}

func InitRepo() (err error) {
	if sqlite, err = sqlitegorm.New(repofile); err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}
	if err = DoAutoMigration(); err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	return
}

func DoAutoMigration() error {
	return sqlite.AutoMigrate(automigration...)
}

func SilentLogger() {
	sqlite.Logger = sqlite.Logger.LogMode(logger.Silent)
}
