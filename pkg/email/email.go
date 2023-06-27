package email

import (
	"net/smtp"
	"os"
)

type EmailResult struct {
	Ok      bool
	Message string
}

func SendEmail(email string, message string) EmailResult {
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

	if err != nil {
		return EmailResult{
			Ok:      false,
			Message: "Failed to send email: " + err.Error(),
		}
	}

	return EmailResult{
		Ok:      true,
		Message: "Email successfully sent",
	}
}
