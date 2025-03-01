package furniture

type modernFurnitureFactory struct{}

// CreateChair implements furniture.FurnitureFactory.
func (m modernFurnitureFactory) CreateChair() Chair {
	return NewModernChair()
}

// CreateSofa implements FurnitureFactory.
func (m modernFurnitureFactory) CreateSofa() Sofa {
	return NewModernSofa()
}

// CreateTable implements FurnitureFactory.
func (m modernFurnitureFactory) CreateTable() Table {
	return NewModernTable()
}

func NewModernFactory() FurnitureFactory {
	return modernFurnitureFactory{}
}
