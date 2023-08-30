package db

import (
	"gorm.io/gorm"
)

type Object struct {
	gorm.Model
	ID      string
	Object  string
	Statues string
}

func CreateObject(id, object string) error {
	tx := db.Create(&Object{ID: id, Object: object, Statues: "created"})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadObject(id string) (*Object, error) {
	obj := &Object{ID: id}
	tx := db.Take(obj)
	if tx.Error != nil {
		return obj, tx.Error
	}
	return obj, nil
}

func UpdateObject(obj *Object) error {
	tx := db.Save(obj)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteObject(id string) error {
	tx := db.Exec("DELETE FROM objects WHERE id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
