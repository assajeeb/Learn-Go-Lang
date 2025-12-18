package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashe, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashe), err

}

func CheckPasswordHash(password, hashpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
}
