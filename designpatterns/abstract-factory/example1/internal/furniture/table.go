package furniture

type modernTable struct{}

func (t modernTable) HasLegs() int {
	return 1
}

func (t modernTable) SitsOn() bool {
	return false
}

func (c modernTable) TypeName() FurnitureType {
	return ModernFurnitureType
}

func NewModernTable() Table {
	return modernTable{}
}

type victorianTable struct{}

func NewVictorianTable() Table {
	return victorianTable{}
}

func (vc victorianTable) HasLegs() int {
	return 4
}

func (vc victorianTable) SitsOn() bool {
	return false
}

func (c victorianTable) TypeName() FurnitureType {
	return VictorianFurnitureType
}
