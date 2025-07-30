package structural

import "testing"

func TestAdapter(t *testing.T) {
	adaptee := &Adaptee{
		Resp: "Adaptee specific response",
	}
	adapter := NewAdapter(adaptee)
	request := adapter.Request()

	if request != adaptee.Resp {
		t.Error("Expected", adaptee.Resp, ", got", request)
	}
}
