package structural

import (
	"testing"
)

func TestBridge(t *testing.T) {
	ascii := &ASCIIRenderer{}
	debug := &DebugRenderer{}

	circle := NewCircle(ascii, 0, 0, 10)
	resp := circle.Draw()
	want := "ASCII: circle at (0.00, 0.00), r=10.00"
	if resp != want {
		t.Error("Circle ASCII draw mismatch. got:", resp, "want:", want)
	}

	circle.SetRenderer(debug)
	resp = circle.Draw()
	want = "DEBUG: CIRCLE{x:0.00,y:0.00,r:10.00}"
	if resp != want {
		t.Error("Debug CIRCLE draw mismatch. got:", resp, "want:", want)
	}

	rectangle := NewRectangle(ascii, 0, 0, 5, 5)
	resp = rectangle.Draw()
	want = "ASCII: rect at (0.00, 0.00), w=5.00, h=5.00"
	if resp != want {
		t.Error("Rectangle ASCII Draw mismatch. got:", resp, "want:", want)
	}

	rectangle.SetRenderer(debug)
	resp = rectangle.Draw()
	want = "DEBUG: RECT{x:0.00,y:0.00,w:5.00,h:5.00}"
	if resp != want {
		t.Error("Debug RECT draw mismatch. got:", resp, "want:", want)
	}
}
