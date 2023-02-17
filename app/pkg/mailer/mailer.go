package mailer

import (
	"github.com/wneessen/go-mail"
)

type Mailer struct {
	client *mail.Client
	cfg    Config
}

type Config struct {
	From     string
	HOST     string
	Port     int
	Username string
	Password string
	SSL      bool
}

func NewMailer(cfg Config) (*Mailer, error) {
	c, err := mail.NewClient(cfg.HOST, mail.WithPort(cfg.Port), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Username), mail.WithPassword(cfg.Password))
	if err != nil {
		return nil, err
	}

	c.SetSSL(cfg.SSL)

	return &Mailer{
		client: c,
		cfg:    cfg,
	}, nil
}

func (c *Mailer) NewMessage(to, subject, body string) error {
	m := mail.NewMsg()
	if err := m.From(c.cfg.From); err != nil {
		return err
	}
	if err := m.To(to); err != nil {
		return err
	}

	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, body)

	if err := c.client.DialAndSend(m); err != nil {
		return nil
	}

	return nil
}
