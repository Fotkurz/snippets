package internal

type PizzaStore struct{}

func NewPizzaStore() PizzaStore {
	return PizzaStore{}
}

func (ps PizzaStore) Order(flavor string) Pizza {
	switch flavor {
	case "pepperoni":
		return NewPepperoniFactory().Prepare()
	case "cheese":
		return NewCheeseFactory().Prepare()
	default:
		return NewRandomPizzaFactory().Prepare()
	}
}
