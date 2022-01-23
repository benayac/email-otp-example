package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/otp-email/middleware"
)

func main() {
	fmt.Println("SERVICE STARTING ON PORT :8000")
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/login_admin", LoginAdminHandler)
	http.HandleFunc("/register", RegisterAccountHandler)
	http.HandleFunc("/validation", ValidateOTPHandler)
	http.HandleFunc("/hash_password", HashPassword)
	http.HandleFunc("/test", middleware.IsAuthorizedUser(TestHandler))
	http.HandleFunc("/test_admin", middleware.IsAuthorizedAdmin(TestHandler))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
