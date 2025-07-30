package structural

type Target interface {
	Request() string
}

type Adaptee struct {
	Resp string
}

func (a *Adaptee) SpecificRequest() string {
	return a.Resp
}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}
