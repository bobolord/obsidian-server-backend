package utilities

import (
	"fmt"
	"log"
	"net/smtp"
)

func NewRegistration(email string) {
	send(email, "Confirmation email", "Please click here to confirm your email address.")
}

func send(receiver string, subject string, body string) {
	fmt.Println("awdawd")
	from := Config.MailerConfig.SenderEmail
	pass := Config.MailerConfig.Password
	to := receiver

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
