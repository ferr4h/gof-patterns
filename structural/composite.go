package structural

type Component interface {
	getPrice() int
}

type Box struct {
	children []Component
}

func NewBox() *Box {
	return &Box{}
}

func (box *Box) getPrice() int {
	price := 0
	for _, child := range box.children {
		price += child.getPrice()
	}
	return price
}

func (box *Box) addChild(component Component) {
	box.children = append(box.children, component)
}

type Product struct {
	price int
}

func NewProduct(price int) *Product {
	return &Product{price}
}

func (product *Product) getPrice() int {
	return product.price
}
