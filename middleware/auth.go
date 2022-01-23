package middleware

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/otp-email/model"
)

var mySigningKey = []byte("unicorns")

func GetJWTUser(client string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = client
	claims["aud"] = "user"
	claims["iss"] = "the.grace"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetJWTAdmin(client string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = client
	claims["aud"] = "admin"
	claims["iss"] = "the.grace"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsAuthorizedUser(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			valid, err := parseAuth(r.Header, "user")
			if err != nil {
				log.Println("[AUTHENTICATION][ERROR] PARSING HEADER: ", err.Error())
				res := model.Response{Status: false, Message: "Authentication Failed"}
				json.NewEncoder(w).Encode(&res)
			}
			if valid {
				endpoint(w, r)
			}
		} else {
			res := model.Response{Status: true, Message: "No Authentication Provided"}
			json.NewEncoder(w).Encode(&res)
		}
	})
}

func IsAuthorizedAdmin(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			valid, err := parseAuth(r.Header, "admin")
			if err != nil {
				log.Println("[AUTHENTICATION][ERROR] PARSING HEADER: ", err.Error())
				res := model.Response{Status: false, Message: "Authentication Failed"}
				json.NewEncoder(w).Encode(&res)
			}
			if valid {
				endpoint(w, r)
			}
		} else {
			res := model.Response{Status: true, Message: "No Authentication Provided"}
			json.NewEncoder(w).Encode(&res)
		}
	})
}

func parseAuth(header map[string][]string, aud string) (bool, error) {
	bearer := strings.Split(header["Authorization"][0], "Bearer ")[1]
	bearer = strings.TrimSpace(bearer)
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid Signing Method")
		}
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			return nil, errors.New("Expired Token")
		}
		checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAudience {
			return nil, errors.New("Invalid aud")
		}
		iss := "the.grace"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return nil, errors.New("Invalid iss")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, nil
	}
	return true, nil
}
