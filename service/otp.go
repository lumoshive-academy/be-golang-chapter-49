package service

import (
	"bytes"
	"fmt"
	"math/rand"
	"text/template"
	"time"

	"gopkg.in/gomail.v2"
)

// Struct untuk memuat data ke dalam template
type EmailData struct {
	OTP string
}

func GenerateOTP() string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	otp := fmt.Sprintf("%06d", rng.Intn(1000000)) // Generate 6 digit OTP
	return otp
}

func SendOTPEmail(to string, otp string) error {
	// Load template HTML
	tmpl, err := template.ParseFiles("email/template.html")
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	// Data yang akan dimuat ke dalam template
	data := EmailData{OTP: otp}

	// Apply template dengan data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	// Buat pesan email
	m := gomail.NewMessage()
	m.SetHeader("From", "youremail@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/html", body.String())

	// Setup SMTP dialer
	d := gomail.NewDialer("smtp.example.com", 587, "youremail@example.com", "yourpassword")

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}
