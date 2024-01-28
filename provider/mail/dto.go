package mail

type SendMailRequest struct {
	To      []string
	Message string
	Subject string
	CC      []string
}
