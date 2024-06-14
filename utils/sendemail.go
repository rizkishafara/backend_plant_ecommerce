package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

func SendForgotPassword(email,paramm string) error {
	subject := "Reset Password"
	body := "Click here to reset your password: http://localhost:3000/resetpassword/" + paramm
	err := SendEmail(email, subject, body)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(500, err.Error())
	}
	return nil
}

func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "plantingpteam200@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, "35660ba17d6714", "79c6356485bce1")

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
