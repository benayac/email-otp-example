package middleware

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

func getRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

func randToken() (string, error) {
	u4, err := uuid.NewV4()
	UUIDtoken := u4.String()
	if err != nil {
		return "", err
	}
	return UUIDtoken, err
}

func GenerateOTP(email string) (string, error) {
	otp, err := getRandNum()
	if err != nil {
		return "", err
	}
	err = SetValue(email, otp, 5*time.Minute)
	if err != nil {
		return "", err
	}
	return otp, nil
}

func ValidateOTP(otp string, email string) (bool, error) {
	originalOtp, err := GetValue(email)
	if err != nil {
		return false, err
	}
	if originalOtp != otp {
		return false, nil
	}
	return true, nil
}
