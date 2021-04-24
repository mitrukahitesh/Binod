package email

import (
	"crypto/tls"
	"net/http"
	"os"

	gomail "gopkg.in/mail.v2"
)

func SendMail(w http.ResponseWriter, r *http.Request, email string, name string) {
	from := "binodcommentor@gmail.com"
	password := os.Getenv("PASSWORD")
	// to := []string{
	// 	email,
	// }

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// Message.
	message := "Dear " + name + ",\n\n" + "Thank you for contacting us.\n\nWe will contact you once you query is processed.\n\nHave a nice day!\n\nRegards\nBinod Youtubewala"

	// // Authentication.
	// auth := smtp.PlainAuth("Binod Youtubewala", from, password, smtpHost)
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Binod Youtubewala Customer Support")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", message)

	// Settings for SMTP server
	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	err := d.DialAndSend(m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
