package utils

import (
	"log"

	"gopkg.in/gomail.v2"
)

type Email struct {
	Email string
	AppPassword string
	Broadcast chan<- []byte
}

type Mail struct {
	From string
	To string
	Subject string
	Body string
}

func NewEmail(e, a string,b chan []byte) *Email {

	return &Email{
		Email:e,
		AppPassword: a,
		Broadcast: b,
	}
}


func (e *Email) SendEmail(mail Mail) {

	m := gomail.NewMessage()
	m.SetHeader("From",mail.From)
	m.SetHeader("To",mail.To)
	m.SetHeader("Subject",mail.Subject)
	m.SetBody("text/html",mail.Body)

	d := gomail.NewDialer("smtp.gmail.com",587,e.Email,e.AppPassword)


	if err := d.DialAndSend(m); err != nil {
		log.Printf("fail to send email: %s",err)
		return
	}
	// e.Broadcast <- []byte("send email successfully")
	
	log.Print("send email successfully")
}