package mydb

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(
	dsn string,
	autoMigrate func(db *gorm.DB) (err error),
) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := autoMigrate(db); err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

// func (db *DB)Where(query any, args ...any) (tx *gorm.DB) {
// 	return db.Where(query, args...)
// }

func (db *DB) Create(tx *gorm.DB, o any) error {
	tx.Create(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (db *DB) Read(tx *gorm.DB, o any) error {
	tx.Where(o).First(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (db *DB) Update(tx *gorm.DB, o any) error {
	tx.Save(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (db *DB) Delete(tx *gorm.DB, o any) error {
	tx.Delete(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (db *DB) MustDelete(o any) {
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

func (db *DB) DB() *gorm.DB {
	return db.db
}

func (db *DB) Begin() *gorm.DB {
	return db.db.Begin()
}

func (db *DB) ErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
