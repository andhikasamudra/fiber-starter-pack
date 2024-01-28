package env

import "os"

func MailFrom() string {
	return os.Getenv("MAIL_FROM")
}

func SMTPHost() string {
	return os.Getenv("SMTP_HOST")
}
func SMTPPort() string {
	return os.Getenv("SMTP_PORT")
}
func SMTPUsername() string {
	return os.Getenv("SMTP_USERNAME")
}
func SMTPPassword() string {
	return os.Getenv("SMTP_PASSWORD")
}
