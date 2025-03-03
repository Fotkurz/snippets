package internal

// Pizza declares the interface for the concrete objects
type Pizza interface {
	Serve() string
}

type PepperoniPizza struct{}

func (pp PepperoniPizza) Serve() string {
	return "serving a pepperoni pizza"
}

type CheesePizza struct{}

func (cp CheesePizza) Serve() string {
	return "serving a cheese piza"
}
