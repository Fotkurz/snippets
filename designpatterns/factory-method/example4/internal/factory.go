package internal

import "math/rand"

type PizzaFactory interface {
	Prepare() Pizza
}

type RandomFactory struct{}

func NewRandomPizzaFactory() PizzaFactory {
	return RandomFactory{}
}

func (rf RandomFactory) Prepare() Pizza {
	if rand.Int()%2 == 0 {
		return PepperoniPizza{}
	} else {
		return CheesePizza{}
	}

}

type PepperoniPizzaFactory struct{}

func NewPepperoniFactory() PizzaFactory {
	return PepperoniPizzaFactory{}
}

func (ppf PepperoniPizzaFactory) Prepare() Pizza {
	return PepperoniPizza{}
}

type CheesePizzaFactory struct{}

func NewCheeseFactory() PizzaFactory {
	return CheesePizzaFactory{}
}

func (cpf CheesePizzaFactory) Prepare() Pizza {
	return CheesePizza{}
}
