package handlers

import (
	"os"
	"encoding/json"
	"crypto/tls"
	"net/http"
	"fmt"

	inp "lol-api/api/v1/input"
	out "lol-api/api/v1/output"

	log "github.com/sirupsen/logrus"

	gomail "gopkg.in/mail.v2"
)

func SendContactMail(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to SendMail route --------")

	var u inp.ContactMail

	if err := json.NewDecoder(_r.Body).Decode(&u); err != nil {
		log.Error("Error decoding body to wishlist struct: ", err)
		out.RespondWithError(_w, out.ErrBodyParams)
		return
	}
	defer _r.Body.Close()

	if err := u.ValidateFields(); err != nil {
		log.Error("Invalid data in wishlist struct: ", err)
		out.RespondWithError(_w, out.ErrBodyParams)
		return
	}

	log.Info(u);

	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("MAIL_USER"))
	m.SetHeader("To", os.Getenv("MAIL_SEND_TO"))

	m.SetHeader("Subject", fmt.Sprintf("Contact us"))
	m.SetBody("text/plain", fmt.Sprintf("NAME: %s\nSENDER: %s\nSUBJECT: %s\nCONTENT: %s", u.Name, u.Sender, u.Subject, u.Content))

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Error("Error sending email ", err)
		out.RespondWithError(_w, err)
		return
	}

	log.Info("-------- Finish SendMail route --------")
	out.RespondWithJSON(_w, http.StatusOK, "Email sent")
}
