package structural

type Pizza interface {
	getPrice() int
}

type Pepperoni struct {
	price int
}

func (p *Pepperoni) getPrice() int {
	return p.price
}

type PizzaWithJalapenos struct {
	basePizza Pizza
}

func (p *PizzaWithJalapenos) getPrice() int {
	return p.basePizza.getPrice() + 15
}

type PizzaWithOlives struct {
	basePizza Pizza
}

func (p *PizzaWithOlives) getPrice() int {
	return p.basePizza.getPrice() + 10
}
