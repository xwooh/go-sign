package notify

import (
	"crypto/tls"
	"github.com/hit9/log"

	gomail "gopkg.in/gomail.v2"
)

var logger = log.Get("notify")

type EmailOpt struct {
	AdminEmail string
	AdminName  string
	Password   string
	SmtpServer string
}

func SendEmail(
	subject, content string,
	emailOpt EmailOpt,
	toEmails ...string,
) {
	m := gomail.NewMessage()

	toEmail := emailOpt.AdminEmail
	if len(toEmails) > 0 {
		toEmail = toEmails[0]
	}

	m.SetAddressHeader("From", emailOpt.AdminEmail, emailOpt.AdminName)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", content)

	smtpHost := "smtp.yeah.net"
	if emailOpt.SmtpServer != "" {
		smtpHost = emailOpt.SmtpServer
	}
	d := gomail.NewDialer(smtpHost, 465, toEmail, emailOpt.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// send email
	if err := d.DialAndSend(m); err != nil {
		logger.Error("send email failed: %v", err)
	} else {
		logger.Info("send email succeed")
	}
}
