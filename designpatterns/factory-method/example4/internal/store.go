package internal

// PizzaStore is used to invoke the correct factory, since
//
//	golang have no inheritance we need to implement it like this
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
