package creational

import "testing"

func TestAbstractFactory(t *testing.T) {
	windowsFactory := GetWidgetFactory("Windows")
	linuxFactory := GetWidgetFactory("Linux")

	windowsButton := windowsFactory.RenderButton()
	linuxButton := linuxFactory.RenderButton()

	if windowsButton.Press() != "WindowsButton is pressed" {
		t.Error("Expected 'WindowsButton is pressed', got ", windowsButton.Press())
	}

	if linuxButton.Press() != "LinuxButton is pressed" {
		t.Error("Expected 'LinuxButton is pressed', got ", linuxButton.Press())
	}

	windowsCheckbox := windowsFactory.RenderCheckbox()
	linuxCheckbox := linuxFactory.RenderCheckbox()

	if windowsCheckbox.Check() != "WindowsCheckbox is checked" {
		t.Error("Expected 'WindowsCheckbox is checked', got ", windowsCheckbox.Check())
	}

	if linuxCheckbox.Check() != "LinuxCheckbox is checked" {
		t.Error("Expected 'LinuxCheckbox is checked', got ", linuxCheckbox.Check())
	}
}
