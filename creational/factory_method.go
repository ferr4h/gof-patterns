package creational

type Notifier interface {
	Send(to, message string) string
}

type EmailNotifier struct {
}

func (notifier EmailNotifier) Send(to, message string) string {
	return "Email to " + to + ": " + message
}

type SMSNotifier struct {
}

func (notifier SMSNotifier) Send(to, message string) string {
	return "SMS to " + to + ": " + message
}

func FactoryMethod(method string) Notifier {
	switch method {
	case "email":
		return EmailNotifier{}
	case "sms":
		return SMSNotifier{}
	}
	return nil
}
