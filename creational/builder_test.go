package creational

import "testing"

func TestBuilder(t *testing.T) {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	product := director.Construct("white", "100 pounds")

	if product.Color != "white" || product.Weight != "100 pounds" {
		t.Error("Product fields do not match")
	}
}
