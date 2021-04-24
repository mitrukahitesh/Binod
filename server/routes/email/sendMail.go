package email

import (
	"net/http"
	"net/smtp"
	"os"
)

func SendMail(w http.ResponseWriter, r *http.Request, email string, name string) {
	from := "binodcommentor@gmail.com"
	password := os.Getenv("PASSWORD")
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Dear " + name + ",\n\n" + "Thank you for contacting us.\n\nWe will contact you once you query is processed.\n\nHave a nice day!\n\nRegards\nBinod Youtubewala")

	// Authentication.
	auth := smtp.PlainAuth("Binod Youtubewala", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
