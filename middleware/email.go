package middleware

import (
	mail "github.com/xhit/go-simple-mail/v2"
)

func defaultSetting(username string, password string) (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 465
	server.Username = username
	server.Password = password
	server.Encryption = mail.EncryptionSSLTLS

	smtpClient, err := server.Connect()
	if err != nil {
		return nil, err
	}
	return smtpClient, err
}

func SendEmail(username string, password string, to string, subject string, body string) error {
	server, err := defaultSetting(username, password)
	if err != nil {
		return err
	}
	defer server.Close()
	email := mail.NewMSG()
	email.SetFrom("From Me <me@host.com>")
	email.AddTo(to)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)
	err = email.Send(server)
	if err != nil {
		return err
	}
	return nil
}
