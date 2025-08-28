package structural

import "testing"

func TestComposite(t *testing.T) {
	product1 := NewProduct(50)
	product2 := NewProduct(100)

	innerBox := NewBox()
	innerBox.addChild(product1)
	innerBox.addChild(product2)

	outerBox := NewBox()
	outerBox.addChild(innerBox)
	outerBox.addChild(product1)

	if outerBox.getPrice() != 200 {
		t.Error("Wrong price")
	}
}
