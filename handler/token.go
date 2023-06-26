package handler

import (
	mailer "email-auth/pkg/email"
	jwt "email-auth/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"net/url"
)

func HandleToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	validateToken := r.URL.Query().Get("validate")
	email := r.URL.Query().Get("email")
	protectedUrl := r.URL.Query().Get("protected-url")

	if len(validateToken) < 1 && len(email) < 1 {
		http.Error(w, "Please provide a token to validate or an email to send a link to", http.StatusBadRequest)
		return
	}

	if len(validateToken) > 1 {
		handleTokenValidation(w, validateToken)
		return
	}

	if len(email) > 1 {
		handleEmailSending(w, email, protectedUrl)
		return
	}

	http.Error(w, "Invalid request", http.StatusBadRequest)
}

func handleTokenValidation(w http.ResponseWriter, validateToken string) {
	res := jwt.ValidateToken(validateToken)

	if !res.Ok {
		http.Error(w, res.Message, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func handleEmailSending(w http.ResponseWriter, email string, protectedUrl string) {
	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "Not a valid email address", http.StatusBadRequest)
		return
	}

	if len(protectedUrl) < 1 {
		http.Error(w, "Please provide protected URL for which the token is issued", http.StatusBadRequest)
		return
	}

	if _, err := url.Parse(protectedUrl); err != nil {
		http.Error(w, "Not a valid email URL", http.StatusBadRequest)
		return
	}

	jwtString, err := jwt.GenerateSignedString()
	if err != nil {
		http.Error(w, "Error generating token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Use this link to login: %s?token=%s", protectedUrl, jwtString)
	mailerResult := mailer.SendEmail(email, message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mailerResult)
}
