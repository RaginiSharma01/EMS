package utils

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendOTPEmail(toEmail, otp string) error {

	from := mail.NewEmail("EMS System", "raginisharma.r07@gmail.com")
	subject := "Your Verification OTP"

	content := fmt.Sprintf(
		"Your OTP for email verification is: %s\n\nThis OTP expires in 5 minutes.",
		otp,
	)

	to := mail.NewEmail("", toEmail)

	message := mail.NewSingleEmail(from, subject, to, content, content)

	client := sendgrid.NewSendClient(os.Getenv("API_KEYS"))

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	fmt.Println("SendGrid Status:", response.StatusCode)

	return nil
}
