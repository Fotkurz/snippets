package internal

import "math/rand"

// The factory only declares the interface
//
//	a impl needs to be considered a PizzaFactory
//	the creation logic for pizza are inside the factories
type PizzaFactory interface {
	Prepare() Pizza
}

// Random PizzaFactory impl
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

// Pepperoni Pizza Factory impl
type PepperoniPizzaFactory struct{}

func NewPepperoniFactory() PizzaFactory {
	return PepperoniPizzaFactory{}
}

func (ppf PepperoniPizzaFactory) Prepare() Pizza {
	return PepperoniPizza{}
}

// Cheese pizza factory impl
type CheesePizzaFactory struct{}

func NewCheeseFactory() PizzaFactory {
	return CheesePizzaFactory{}
}

func (cpf CheesePizzaFactory) Prepare() Pizza {
	return CheesePizza{}
}
