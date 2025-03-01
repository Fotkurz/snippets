package furniture

type Furniture interface {
	TypeName() FurnitureType
	HasLegs() int
	SitsOn() bool
}

type (
	Chair interface {
		Furniture
	}

	Table interface {
		Furniture
	}

	Sofa interface {
		Furniture
	}
)
