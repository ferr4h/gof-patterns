package creational

import "testing"

func TestFactoryMethod(t *testing.T) {
	emailNotifier := FactoryMethod("email")
	smsNotifier := FactoryMethod("sms")

	expectedEmail := "Email to user@site.com: Hello!"
	expectedSMS := "SMS to +12345: Hi!"

	if emailNotifier.Send("user@site.com", "Hello!") != expectedEmail {
		t.Error("Expected: " + expectedEmail + ", got: " + emailNotifier.Send("user@site.com", "Hello!"))
	}

	if smsNotifier.Send("+12345", "Hi!") != expectedSMS {
		t.Error("Expected: " + expectedSMS + ", got: " + SMSNotifier{}.Send("+12345", "Hi!"))
	}
}
