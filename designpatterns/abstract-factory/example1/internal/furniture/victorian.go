package furniture

type victorianFurnitureFactory struct {
}

// CreateChair implements FurnitureFactory.
func (v victorianFurnitureFactory) CreateChair() Chair {
	return NewVictorianChair()
}

// CreateSofa implements FurnitureFactory.
func (v victorianFurnitureFactory) CreateSofa() Sofa {
	return NewVictorianSofa()
}

// CreateTable implements FurnitureFactory.
func (v victorianFurnitureFactory) CreateTable() Table {
	return NewVictorianTable()
}

func NewVictorianFurnitureFactory() FurnitureFactory {
	return victorianFurnitureFactory{}
}
