package utils

import (
	"fmt"
	"net/smtp"
)

func SendVerificationEmail(toEmail, verificationLink, fromEmail, password string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", fromEmail, password, smtpHost)

	body := fmt.Sprintf(
		"Subject: Verify your email\r\n\r\nClick the link below to verify your email:\r\n\r\n%s\r\n\r\nThis link expires in 10 minutes.",
		verificationLink,
	)

	return smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		fromEmail,
		[]string{toEmail},
		[]byte(body),
	)
}
