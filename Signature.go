package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

func GenKeys() (*rsa.PrivateKey, rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic (err)
	}
	pubkey := privkey.PublicKey

	return privkey, pubkey
}