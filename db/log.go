package db

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Header *string
	Body   *string
	Verify *string
}

func CreateLog(header, body, verify *string) error {
	tx := db.Create(&Log{Header: header, Body: body, Verify: verify})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadLog(id uint) (*Log, error) {
	log := Log{}
	tx := db.Raw("SELECT * FROM logs WHERE id = ?", id).Scan(&log)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &log, nil
}

func UpdateLog(obj *Log) error {
	tx := db.Save(obj)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteLog(id uint) error {
	tx := db.Exec("DELETE FROM logs WHERE id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
