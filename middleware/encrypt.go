package middleware

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pass []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func CompareHashAndPassword(hashedPass string, plainPass []byte) (bool, error) {
	byteHash := []byte(hashedPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil {
		return false, err
	}
	return true, nil
}
