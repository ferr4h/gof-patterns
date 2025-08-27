package structural

import "testing"

func TestProxy(t *testing.T) {
	app := Application{}

	proxy := NewProxy(&app, 1)
	code, _ := proxy.HandleRequest("/notauser", "GET")
	want := 404
	if code != want {
		t.Errorf("got %d, want %d", code, want)
	}

	code, _ = proxy.HandleRequest("/user", "POST")
	want = 405
	if code != want {
		t.Errorf("got %d, want %d", code, want)
	}

	code, _ = proxy.HandleRequest("/user", "GET")
	want = 403
	if code != want {
		t.Errorf("got %d, want %d", code, want)
	}
}
