package dao

import (
	"errors"
	"fmt"
	"time"

	core "github.com/Hana-ame/fedi-antenna/core/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	core.AutoMigrate(db)
}

// func Where(query any, args ...any) (tx *gorm.DB) {
// 	return db.Where(query, args...)
// }

func Create(tx *gorm.DB, o any) error {
	tx.Create(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Read(tx *gorm.DB, o any) error {
	tx.Where(o).First(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Update(tx *gorm.DB, o any) error {
	tx.Save(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Delete(tx *gorm.DB, o any) error {
	tx.Delete(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func MustDelete(o any) {
	tx := db.Begin()
	for i := 0; i < 10; i++ {
		tx.Delete(o)
		time.Sleep(time.Second * time.Duration(i))
		if tx.Error == nil {
			return
		}
	}
	fmt.Printf("not deleted. %s\n", o)
}

func DB() *gorm.DB {
	return db
}

func Begin() *gorm.DB {
	return db.Begin()
}

func ErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
