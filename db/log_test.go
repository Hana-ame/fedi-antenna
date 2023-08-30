package db

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateLog(t *testing.T) {
	Init()
	s := "12321"
	err := CreateLog(&s, &s, &s)
	if err != nil {
		t.Error(err)
	}
}

func TestReadLog(t *testing.T) {
	Init()
	a, err := ReadLog(1)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(a)
	b, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", b)
}

func TestDeleteLog(t *testing.T) {
	Init()
	err := DeleteLog(1)
	if err != nil {
		t.Error(err)
	}
}
