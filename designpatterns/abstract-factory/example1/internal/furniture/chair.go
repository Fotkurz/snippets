package furniture

type modernChair struct{}

func (c modernChair) HasLegs() int {
	return 1
}

func (c modernChair) SitsOn() bool {
	return true
}

func (c modernChair) TypeName() FurnitureType {
	return ModernFurnitureType
}

func NewModernChair() Chair {
	return modernChair{}
}

type victorianChair struct{}

func NewVictorianChair() Chair {
	return victorianChair{}
}

func (vc victorianChair) HasLegs() int {
	return 4
}

func (vc victorianChair) SitsOn() bool {
	return true
}

func (c victorianChair) TypeName() FurnitureType {
	return VictorianFurnitureType
}
