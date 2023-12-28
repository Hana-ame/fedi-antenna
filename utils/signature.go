package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"strconv"
)

var RSABits, _ = strconv.Atoi(os.Getenv("RSA_BITS"))

// create a new private key
func GeneratePrivateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, RSABits)
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	return privateKey
}

// marshal private key to pem format
func MarshalPrivateKey(privateKey *rsa.PrivateKey) string {
	bytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	privatePem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: bytes,
		},
	)
	return string(privatePem)
}

// marshal public key to pem format
func MarshalPublicKey(publicKey *rsa.PublicKey) string {
	bytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	publicPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: bytes,
		},
	)
	return string(publicPem)
}
