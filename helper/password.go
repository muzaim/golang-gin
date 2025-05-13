package helper

import "golang.org/x/crypto/bcrypt"

func HashPasword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHash), err
}
