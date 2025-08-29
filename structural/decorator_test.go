package structural

import "testing"

func TestDecorator(t *testing.T) {
	pepperoni := &Pepperoni{100}
	pepperoniWithJalapenos := &PizzaWithJalapenos{pepperoni}
	pepperoniWithJalapenosAndOlives := &PizzaWithOlives{pepperoniWithJalapenos}

	if pepperoniWithJalapenosAndOlives.getPrice() != 125 {
		t.Error("Incorrect price")
	}
}
