package mail

type Interface interface {
	Send(param SendMailRequest) error
}
