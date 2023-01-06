package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var p = &Params{
	memory:      64 * 1024, // 64 MB
	iterations:  3,
	parallelism: 1,
	saltLength:  16,
	keyLength:   32,
}

func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// func decodeHash(encodedHash string,salt string){

// }

func GenerateHashfromPassword(password string) (string, string, error) {
	salt, err := GenerateRandomBytes(p.saltLength)
	if err != nil {
		return "", "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// fmt.Printf("HashedPassword :%v", hash)

	base64Hash := base64.RawStdEncoding.EncodeToString(hash)
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)

	return base64Hash, base64Salt, nil

}

func VerifyPassword(password, encodedHash, salt string) (bool, error) {

	Hash, _ := base64.RawStdEncoding.Strict().DecodeString(encodedHash)
	Salt, _ := base64.RawStdEncoding.Strict().DecodeString(salt)
	passwordHash := argon2.IDKey([]byte(password), Salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	T := base64.RawStdEncoding.EncodeToString(passwordHash)
	fmt.Printf("INPUT: %v\n INPUT_HASH:%s\n DBPASS: %s\n", password, T, encodedHash)
	if subtle.ConstantTimeCompare(passwordHash, Hash) == 1 {
		return true, nil
	}
	
	return false, nil
}
