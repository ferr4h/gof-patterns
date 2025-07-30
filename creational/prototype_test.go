package creational

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	orig := &Document{
		Title: "Title",
		Metadata: map[string]string{
			"author": "John Doe",
		},
	}
	cloneProto := orig.Clone()
	clone, _ := cloneProto.(*Document)

	if clone == orig {
		t.Error("A new object was expected, but the pointers match")
	}

	if clone.Title != orig.Title || clone.Metadata["author"] != orig.Metadata["author"] {
		t.Error("Fields do not match")
	}
}
