package mailernats

// Mailer is
type Mailer interface {
	Send(data *EmailData) error
}
