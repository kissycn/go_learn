package mail

type Mailer interface {
	SendMail(subject, dest, body string) error
}

type MailClient struct {
	m Mailer
}

func New(m Mailer) *MailClient {
	return &MailClient{
		m: m,
	}
}

func (c *MailClient) ComposeAndSend(subject, dest, body string) error {
	err := c.m.SendMail(subject, dest, body)
	if nil != err {
		return err
	}
	return nil
}
