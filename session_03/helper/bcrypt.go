package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	length := 12
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, length)

	return string(hash)
}

func ComparePassword(h, p []byte) bool {
	hash, password := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
