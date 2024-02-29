package dao

import (
	"crypto/rsa"
	"log"

	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// the most direct way to read local user's privateKey
func ReadPrivateKeyByOwner(id string) (pk *rsa.PrivateKey, err error) {
	// tx := db.Begin()

	lu := &core.LocalUser{
		ActivitypubID: id,
	}
	err = Read(db, lu)
	if err != nil {
		log.Println(err)
		return
	}
	pk, err = utils.ParsePrivateKey(lu.PrivateKeyPem)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
