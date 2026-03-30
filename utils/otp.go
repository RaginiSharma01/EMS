package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// init runs automatically when the package loads
// It seeds the random number generator once
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateOTP generates a 6-digit numeric OTP
func GenerateOTP() string {
	otp := rand.Intn(900000) + 100000 // ensures 6 digits
	return fmt.Sprintf("%d", otp)
}
