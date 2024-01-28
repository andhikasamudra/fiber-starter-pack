package mail

import (
	"crypto/tls"
	"github.com/andhikasamudra/fiber-starter-pack/internal/env"
	simpleMail "github.com/xhit/go-simple-mail/v2"
	"strconv"
	"time"
)

func (p *Provider) Send(param SendMailRequest) error {
	server := simpleMail.NewSMTPClient()
	server.Host = env.SMTPHost()
	server.Port, _ = strconv.Atoi(env.SMTPPort())
	server.Username = env.SMTPUsername()
	server.Password = env.SMTPPassword()
	server.Encryption = simpleMail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	conn, err := server.Connect()
	if err != nil {
		p.Logger.Error("failed to connect to mail provider")
		return err
	}

	email := simpleMail.NewMSG()
	email.SetFrom(env.MailFrom()).
		AddTo(param.To...).
		SetSubject(param.Subject)

	if len(param.CC) > 0 {
		email.AddCc(param.CC...)
	}
	email.SetBody(simpleMail.TextHTML, param.Message)

	err = email.Send(conn)
	if err != nil {
		return err
	}

	return nil
}
