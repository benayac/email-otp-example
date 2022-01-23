package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/otp-email/database"
	"github.com/otp-email/middleware"
	"github.com/otp-email/model"
)

func RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(("[REGISTER ACCOUNT][REQUEST]"))
	decoder := json.NewDecoder(r.Body)
	var req model.RegisterRequest
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("[REGISTER ACCOUNT][ERROR] DECODE REQUEST:" + err.Error())
	}
	hashedPass, err := middleware.HashAndSalt([]byte(req.Password))
	if err != nil {
		log.Println("[REGISTER ACCOUNT][ERROR] FAIL TO HASH PASSWORD:" + err.Error())
	}
	req.Password = hashedPass
	err = database.InsertUser(&req)
	if err != nil {
		log.Println("[REGISTER ACCOUNT][ERROR] INSERT TO DB:" + err.Error())
	}
	otp, err := middleware.GenerateOTP(req.Email)
	if err != nil {
		log.Println("[REGISTER ACCOUNT][ERROR] GENERATE OTP:" + err.Error())
	}
	go func() {
		err = middleware.SendEmail("emailemail75@gmail.com", "vwabvxacdzxiizeo", req.Email, "OTP", otp)
		if err != nil {
			log.Println("[REGISTER ACCOUNT][ERROR] SENDING EMAIL:" + err.Error())
		}
		log.Println(("[REGISTER ACCOUNT][EMAIL SENT]"))
	}()
	res := model.Response{Status: true, Message: "Register Success"}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func ValidateOTPHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(("[VALIDATE PASSCODE][REQUEST]"))
	decoder := json.NewDecoder(r.Body)
	var otp model.OTPRequest
	err := decoder.Decode(&otp)
	if err != nil {
		log.Println("[VALIDATE PASSCODE][ERROR] DECODE REQUEST:" + err.Error())
	}
	result, err := middleware.ValidateOTP(otp.OTP, otp.Email)
	if err != nil {
		log.Println("[VALIDATE PASSCODE][ERROR] VALIDATE OTP:" + err.Error())
	}
	var res model.OTPResponse
	if result {
		res.Status = true
		err = database.UpdateUserVerification(otp.Email)
		if err != nil {
			log.Println("[VALIDATE PASSCODE][ERROR] UPDATE VERIFICATION STATUS:" + err.Error())
		}
	} else {
		res.Status = false
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(("[LOGIN][REQUEST]"))
	decoder := json.NewDecoder(r.Body)
	var login model.LoginRequest
	err := decoder.Decode(&login)
	if err != nil {
		log.Println("[LOGIN][ERROR] DECODE REQUEST " + err.Error())
	}
	hashedPassword, err := database.SelectPasswordUser(login.Email)
	if err != nil {
		log.Println("[LOGIN][ERROR] SELECT PASSWORD " + err.Error())
	}
	validation, err := middleware.CompareHashAndPassword(hashedPassword, []byte(login.Password))
	if err != nil {
		log.Println("[LOGIN][ERROR] COMPARE HASH PASSWORD " + err.Error())
	}
	var res model.Response
	if validation {
		jwt, err := middleware.GetJWTUser(login.Email)
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

func HashPassword(w http.ResponseWriter, r *http.Request) {
	password := r.URL.Query().Get("password")
	hashedPass, err := middleware.HashAndSalt([]byte(password))
	if err != nil {
		log.Println("[HASH PASSWORD][ERROR] GET HASH")
	}
	res := model.Response{Status: true, Message: hashedPass}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	res := model.Response{Status: true, Message: "Success"}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&res)
}
