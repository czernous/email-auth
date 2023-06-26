package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type EmailResult struct {
	Ok      bool
	Message string
}

func getResultMessage(e error) string {
	if e != nil {
		return "Failed to send email: " + e.Error()
	}

	return "Email successfully sent"
}

func loadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}
}

// TODO: change to use environment variables
func SendEmail(email string, message string) EmailResult {
	loadEnv()

	port := os.Getenv("SMTP_PORT")
	host := os.Getenv("SMTP_HOST")
	senderEmail := os.Getenv("SMTP_LOGIN")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", senderEmail, password, host)

	to := []string{email}

	msg := []byte("To: " + email + "\r\n" +

		"From: " + senderEmail + "\n" +

		"Subject: Authentication link\r\n" +

		"\r\n" +

		message + "\r\n")

	err := smtp.SendMail(host+":"+port, auth, senderEmail, to, msg)

	res := EmailResult{
		Ok:      err != nil,
		Message: getResultMessage(err),
	}

	return res
}
