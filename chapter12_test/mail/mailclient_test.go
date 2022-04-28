package mail

import "testing"

type fakeMailer struct{}

func (f fakeMailer) SendMail(subject, dest, body string) error {
	return nil
}

func TestMailClient_ComposeAndSend(t *testing.T) {
	m := &fakeMailer{}
	mc := New(m)
	err := mc.ComposeAndSend("s", "d", "b")
	if err != nil {
		t.Errorf("want nil , got %v", err)
	}
}
