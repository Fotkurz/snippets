package furniture

type modernSofa struct{}

func (s modernSofa) HasLegs() int {
	return 1
}

func (s modernSofa) SitsOn() bool {
	return true
}

func (c modernSofa) TypeName() FurnitureType {
	return ModernFurnitureType
}

func NewModernSofa() Sofa {
	return modernSofa{}
}

type victorianSofa struct{}

func NewVictorianSofa() Sofa {
	return victorianSofa{}
}

func (vc victorianSofa) HasLegs() int {
	return 4
}

func (vc victorianSofa) SitsOn() bool {
	return true
}

func (c victorianSofa) TypeName() FurnitureType {
	return VictorianFurnitureType
}
