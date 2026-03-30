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

	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<body style="font-family:Arial;background:#f4f6f8;padding:40px">
<div style="max-width:500px;margin:auto;background:white;padding:30px;border-radius:8px">

<h2 style="text-align:center;color:#2c3e50">Employee Management System</h2>

<p>Hello,</p>

<p>Use the OTP below to verify your email:</p>

<div style="text-align:center;font-size:28px;font-weight:bold;letter-spacing:5px;color:#1a73e8">
%s
</div>

<p>This OTP will expire in <b>5 minutes</b>.</p>

</div>
</body>
</html>
`, otp)

	to := mail.NewEmail("", toEmail)

	message := mail.NewSingleEmail(
		from,
		subject,
		to,
		"Your OTP is: "+otp,
		htmlContent,
	)

	client := sendgrid.NewSendClient(os.Getenv("API_KEYS"))

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	fmt.Println("SendGrid Status:", response.StatusCode)

	return nil
}
