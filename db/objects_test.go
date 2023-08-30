package db

import (
	"fmt"
	"testing"

	"github.com/Hana-ame/fedi-antenna/utils"
)

// pass
func TestObjectCreate(t *testing.T) {
	Init()
	id := utils.NewUUIDString()
	err := CreateObject(id, "id"+id+"id")
	if err != nil {
		t.Error(err)
	}
}
func TestObjectRead(t *testing.T) {
	Init()
	id := "facf97b0-77be-4b10-bb0c-ac6975a81bc1"
	a, err := ReadObject(id)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)
}
func TestObjectUpdate(t *testing.T) {
	Init()
	id := "facf97b0-77be-4b10-bb0c-ac6975a81bc1"
	a, err := ReadObject(id)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)
	a.Statues = "edited"
	err = UpdateObject(a)
	if err != nil {
		t.Error(err)
	}
}
func TestObjectDelete(t *testing.T) {
	Init()
	id := "facf97b0-77be-4b10-bb0c-ac6975a81bc1"
	err := DeleteObject(id)
	if err != nil {
		t.Error(err)
	}
}
