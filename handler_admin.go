package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/otp-email/database"
	"github.com/otp-email/middleware"
	"github.com/otp-email/model"
)

func LoginAdminHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(("[LOGIN ADMIN][REQUEST]"))
	decoder := json.NewDecoder(r.Body)
	var login model.LoginRequest
	err := decoder.Decode(&login)
	if err != nil {
		log.Println("[LOGIN ADMIN][ERROR] DECODE REQUEST " + err.Error())
	}
	hashedPassword, err := database.SelectPasswordAdmin(login.Email)
	if err != nil {
		log.Println("[LOGIN ADMIN][ERROR] SELECT PASSWORD " + err.Error())
	}
	validation, err := middleware.CompareHashAndPassword(hashedPassword, []byte(login.Password))
	if err != nil {
		log.Println("[LOGIN ADMIN][ERROR] COMPARE HASH PASSWORD " + err.Error())
	}
	var res model.Response
	if validation {
		jwt, err := middleware.GetJWTAdmin(login.Email)
		if err != nil {
			log.Println("[LOGIN][ERROR] GET JWT AUTH" + err.Error())
		}
		res.Status = true
		res.Message = jwt
	} else {
		res.Status = true
		res.Message = "Login Failed"
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&res)
}
