package creational

type Product struct {
	Color  string
	Weight string
}

type Builder interface {
	SetColor(string)
	SetWeight(string)
	Build() *Product
}

type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &Product{}}
}

func (b *ConcreteBuilder) SetColor(color string) {
	b.product.Color = color
}

func (b *ConcreteBuilder) SetWeight(weight string) {
	b.product.Weight = weight
}

func (b *ConcreteBuilder) Build() *Product {
	return b.product
}

type Director struct {
	builder *ConcreteBuilder
}

func NewDirector(builder *ConcreteBuilder) *Director {
	return &Director{builder}
}

func (d *Director) SetBuilder(builder *ConcreteBuilder) {
	d.builder = builder
}

func (d *Director) Construct(color, weight string) *Product {
	d.builder.SetColor(color)
	d.builder.SetWeight(weight)
	return d.builder.Build()
}
