package creational

type WidgetFactory interface {
	RenderButton() Button
	RenderCheckbox() Checkbox
}

type Button interface {
	Press() string
}

type Checkbox interface {
	Check() string
}

func GetWidgetFactory(system string) WidgetFactory {
	if system == "Windows" {
		return &WindowsFactory{}
	}
	if system == "Linux" {
		return &LinuxFactory{}
	}

	return nil
}

type WindowsFactory struct {
}

type WindowsButton struct {
}

func (w WindowsButton) Press() string {
	return "WindowsButton is pressed"
}

type WindowsCheckbox struct {
}

func (w WindowsCheckbox) Check() string {
	return "WindowsCheckbox is checked"
}

func (factory *WindowsFactory) RenderButton() Button {
	return &WindowsButton{}
}

func (factory *WindowsFactory) RenderCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

type LinuxFactory struct {
}

type LinuxButton struct {
}

func (linux LinuxButton) Press() string {
	return "LinuxButton is pressed"
}

type LinuxCheckbox struct {
}

func (linux LinuxCheckbox) Check() string {
	return "LinuxCheckbox is checked"
}

func (factory *LinuxFactory) RenderButton() Button {
	return &LinuxButton{}
}

func (factory *LinuxFactory) RenderCheckbox() Checkbox {
	return &LinuxCheckbox{}
}
